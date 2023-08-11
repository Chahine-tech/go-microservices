# Use an official Go runtime as the base image
FROM golang:1.20.1-alpine3.13

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and the application code
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Build the application
RUN go build -o main ./cmd

# Expose the port the application runs on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
