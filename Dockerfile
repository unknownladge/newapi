# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.13.4-alpine3.10

# Add Maintainer Info
LABEL maintainer="Pasit <pasit>"

# Set the Current Working Directory inside the container
WORKDIR /newapi

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 10000

# Command to run the executable
CMD ["./main","run","--host","0.0.0.0"]