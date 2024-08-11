# Use official Go image as the base
FROM golang:1.22.4 as builder

# Set the author label
LABEL authors="Luis Sosa"

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files for dependency installation
COPY go.mod go.sum ./

# Install any needed dependencies specified in go.mod
RUN go mod download

# Copy the rest of the application's source code
COPY . .

# Build the Go application
RUN go build -o lexilearn main.go

# Use a minimal image for the final build
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the pre-built binary from the builder
COPY --from=builder /app/lexilearn .

# Run the Go application
CMD ["./lexilearn"]
