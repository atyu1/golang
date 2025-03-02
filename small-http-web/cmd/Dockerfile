# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code
COPY main.go main.go

# Build the Go binary
RUN GOARCH=arm GOOS=linux GOARM=7 go build -o server main.go

# Stage 2: Create a small image with the binary
FROM scratch

# Copy the binary from the builder stage
COPY --from=builder /app/server /server

# Expose the port that the server will run on
EXPOSE 8080

# Command to run the binary
ENTRYPOINT ["/server"]