package main

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
  "strings"
)

func handleListEnv(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  env := make(map[string]string)

  for _, pair := range os.Environ() {
    l := strings.Split(pair, "=")
    env[l[0]] = l[1]
  }
  result, err := json.Marshal(env)
  if err != nil {
    http.Error(w, "Error preparing the env", 500)
    return
  }
  w.Write(result)
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
  who := r.URL.Query().Get("who")
  if who == "" {
    fmt.Fprintf(w, "Hello World!")
    return
  }

  fmt.Fprintf(w, "Hello %s!!", who)
}

func main() {
  http.HandleFunc("/env", handleListEnv)
  http.HandleFunc("/welcome", handleWelcome)
  
  port := os.Getenv("LISTENING_PORT")
  if port == "" {
    port = "80"
  }
  addr := fmt.Sprintf(":%s", port)
  log.Printf("Starting server at %s", addr)

  err := http.ListenAndServe(addr, nil)
  if err != nil {
    log.Fatalf("Failed to start http server:%v", err)
  }
}
