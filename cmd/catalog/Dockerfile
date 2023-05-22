# Use the official Golang image as the base image
FROM golang:1.19-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o catalog-service .

# Use a minimal base image for the final image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final image
COPY --from=builder /app/catalog-service .

# Set any necessary environment variables
ENV PORT=8080

# Expose the port on which the service will listen
EXPOSE $PORT

# Run the Go application
CMD ["./catalog-service"]
