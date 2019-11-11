# Dockerfile References: https://docs.docker.com/engine/reference/builder/
#IMPORTANT : in some case localhost:8080 isn't accessible, instead the app runs and accessible on docker-machine ip
#run command on terminal : docker-machine ip
#example : http://192.168.99.100:8080/vocalconsonant/osama

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="gimindika@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]