# Use an official Golang runtime as a parent image
FROM golang:latest

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Download and install any required dependencies
RUN go mod download

WORKDIR /app/cmd/ab-backend

# Build the Go app
RUN go build 

# Expose port 3000 for incoming traffic
EXPOSE 3000

# Define the command to run the app when the container starts
CMD ["./ab-backend","-stage","production"]