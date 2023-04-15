# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine AS builder

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the current directory contents into the container at /go/src/app
COPY . .

# Build the Go app with optimizations enabled
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main .

# Use a minimal base image for the final container
FROM alpine:latest

# Add CA certificates for SSL/TLS support
RUN apk --no-cache add ca-certificates

# Expose port 8080 for the API server
EXPOSE 8080

# Set the working directory to /app
WORKDIR /app

# Copy the compiled binary from the builder image to the final image
COPY --from=builder /go/src/app/main .

# Run the command to start the API server
CMD ["./main"]