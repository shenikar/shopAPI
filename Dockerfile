# --- Build Stage ---
FROM golang:1.24-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# -o ./out/shop-api: specifies the output file name and location.
# -ldflags="-w -s": strips debugging information, reducing the binary size.
# ./cmd/shopApi: is the path to the main package.
RUN CGO_ENABLED=0 GOOS=linux go build -o ./out/shop-api -ldflags="-w -s" ./cmd/shopApi

# --- Deploy Stage ---
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/out/shop-api .

# Copy migrations
COPY migrations ./migrations

# Copy .env_example as .env, in a real scenario this should be handled by CI/CD or orchestration tools
COPY .env_example .env

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./shop-api"]