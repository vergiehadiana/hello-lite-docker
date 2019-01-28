FROM alpine
RUN apk update && apk add ca-certificates curl && rm -rf /var/cache/apk

WORKDIR /hello
ADD hello .

CMD ["./hello"]
