# start from base
FROM golang:1.11-alpine
MAINTAINER Vergie Hadiana <vergie.gmail.com>

# install system-wide
RUN apk update && apk add ca-certificates curl && rm -rf /var/cache/apk

# copy our application code
ADD . /opt/hello
WORKDIR /opt/hello

# fetch app specific deps
RUN export CGO_ENABLED=0
RUN go build -o hello main.go

# expose port
EXPOSE 8080

# start app
CMD ["./hello"]
