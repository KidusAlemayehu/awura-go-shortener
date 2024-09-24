# Start with a base image containing Go
FROM golang:1.18-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all necessary Go modules
RUN go mod download

# Build the Go app
RUN go build -o main cmd/main.go

# Expose port 18080 to the outside world
EXPOSE 18080

# Command to run the executable
CMD ["./main"]
