# Builder stage
FROM golang:latest AS builder

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Create and set the working directory
WORKDIR /app

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application files and build the app
COPY . .
RUN go build -o app .

# Final stage: use a lightweight scratch image for production
FROM alpine:latest

# Install any necessary CA certificates (for connecting securely to external services)
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /root/

# Copy the compiled app from the builder stage
COPY --from=builder /app/app .

# Expose the port
EXPOSE 8080

# Set the entry point to the executable
CMD ["./app"]