# Use the official Golang image as a parent image
FROM golang:1.17 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and download them
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o MyFirstGoApi .

# Use a lightweight alpine image for the runtime environment
FROM alpine:latest

# Add the necessary certificates for net requests
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /root/

# Copy the binary from the build stage
COPY --from=builder /app/MyFirstGoApi .

# Expose the port your app runs on
EXPOSE 8080

# Command to run the application
CMD ["./MyFirstGoApi"]
