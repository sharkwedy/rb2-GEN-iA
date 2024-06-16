# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
FROM golang:1.22.4 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . .

# Build the binary.
# -o myapp specifies the output file name, which will be placed in the current directory.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
FROM debian:buster-slim
WORKDIR /root/
# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp ./

# Run the web service on container startup.
CMD ["./myapp"]