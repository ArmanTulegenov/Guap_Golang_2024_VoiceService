# Use the official Golang image as the base image
FROM golang:1.22.0

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY ./cmd ./cmd
COPY ./web/public ./public
RUN mkdir -p /app/local_dir
COPY go.mod ./
COPY go.sum ./

# Build the Go application
RUN go build -o app ./cmd/voice_service 

# Set the entry point for the container
CMD ["./app"]
