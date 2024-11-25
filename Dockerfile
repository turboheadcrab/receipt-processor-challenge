# Stage 1: Build the Go application
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o receipt-processor

# Stage 2: Runtime container
FROM debian:bookworm-slim

# Set the working directory
WORKDIR /app

# Install certificates for HTTPS
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the built binary from the build stage
COPY --from=build /app/receipt-processor .

# Expose the port the application will listen on
EXPOSE 8181

# Command to run the application
CMD ["/app/receipt-processor"]
