FROM golang:latest as builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o motion-go src/cmd/main.go


FROM scratch
COPY --from=builder /app/motion-go /motion-go
COPY config.properties /config.properties
ENTRYPOINT ["/motion-go"]