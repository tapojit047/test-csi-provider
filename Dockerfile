# Use the official Golang image as the base image
FROM golang:1.23 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o test-csi-server main.go

# Use a minimal image for the runtime
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/test-csi-server .

# Ensure the binary is executable
RUN chmod +x /app/test-csi-server

# Expose a Unix socket for the CSI driver communication
CMD ["./test-csi-server"]
