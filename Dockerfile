FROM golang:1.9

COPY ./src /app/src
ENV GOPATH "/go:/app"
RUN go build -o /main /app/src/main.go

ENTRYPOINT /main
