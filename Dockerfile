FROM zhihaojun/golang-web-scaffold

COPY ./src /app/src
ENV GOPATH "/go:/app"
RUN go build -o /main /app/src/main.go
EXPOSE 1323

ENTRYPOINT /main
