# Use the official Go image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go project files into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Run the Go application
CMD ["./main"]
