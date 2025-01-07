# Stage 1: Build the Go application
FROM golang:1.19.2-bullseye AS builder

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o api .

# Stage 2: Create a lightweight image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/api .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./api"]