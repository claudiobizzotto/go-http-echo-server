FROM golang:alpine3.8
COPY echoserver.go /
EXPOSE 8080
WORKDIR /
ENTRYPOINT ["go", "run", "echoserver.go", "8080"]