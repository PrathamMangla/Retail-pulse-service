# Stage 1: Build
FROM golang:1.23.4 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest

# Set the working directory in the container
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/main .

# Expose the port the app will run on
EXPOSE 8080

# Run the application
CMD ["./main"]