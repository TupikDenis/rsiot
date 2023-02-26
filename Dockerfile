FROM golang:latest AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ./go-demo ./cmd/web/.

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/go-demo .
EXPOSE 8080
CMD ["/go-demo"]