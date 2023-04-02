FROM golang:1.19-alpine as builder

WORKDIR /app
COPY . .
COPY go.mod .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o server ./cmd/main.go

FROM scratch
COPY --from=builder /app/server .
COPY --from=builder /app/config.properties .
CMD ["./server"]