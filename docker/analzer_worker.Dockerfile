# Stage 1: Build
FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

RUN go build -o /app/analzer_worker cmd/worker/main.go

# Stage 2: Run
FROM alpine:latest AS runner
WORKDIR /app
COPY --from=builder /app/analzer_worker .
CMD ["./analzer_worker"]

