# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .

# Build the Go application
RUN go build -v -o main .

# Expose the port the application runs on
EXPOSE 8080

# Command to run the executable
RUN chmod +x ./main
CMD ["./main"]
