# 1. Build stage
FROM golang:1.24.1-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go module and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
# CGO_ENABLED=0 builds a static binary, which is what we want for a distroless container
# -ldflags="-w -s" strips debug information, reducing the binary size
RUN CGO_ENABLED=0 go build -o /go/bin/app -ldflags="-w -s" ./cmd/api

# 2. Production stage
FROM gcr.io/distroless/static-debian11 AS prod

# Set environment variable for production config
ENV APP_ENV=production

# Copy the static binary from the builder stage
COPY --from=builder /go/bin/app /app

# Copy the production config from the builder stage
COPY --from=builder /app/config/config-prod.yaml /config/config-prod.yaml

# Expose port 8081 to the outside world
EXPOSE 8081

# Command to run the executable
ENTRYPOINT ["/app"]