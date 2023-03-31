FROM golang:latest

RUN mkdir "/app"
WORKDIR "/app"

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go", "run", "cmd/main.go"]
