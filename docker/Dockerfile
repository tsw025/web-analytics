# Stage 1: Build the application
FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Install curl and download the migrate tool
# Install curl, download and extract the migrate tool
RUN apk add curl \
    && curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar -xz -C /usr/local/bin/ \
    && chmod +x /usr/local/bin/migrate


RUN chmod +x scripts/entrypoint.sh
RUN chmod +x scripts/service.sh

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go


# Stage 2: Run the application
FROM alpine:latest as deploy

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

# Copy the binary from the build stage
COPY --from=builder /app/server .
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate
COPY --from=builder /app/scripts/entrypoint.sh /entrypoint.sh
COPY --from=builder /app/scripts/service.sh /service.sh
EXPOSE 8080

ENTRYPOINT /entrypoint.sh $0 $@
CMD ["/service.sh"]
