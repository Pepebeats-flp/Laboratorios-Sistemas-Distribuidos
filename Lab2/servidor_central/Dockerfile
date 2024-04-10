# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# WORKDIR /go/src/app

FROM golang:1.22.2

# Create a directory for the app
RUN mkdir /app
# Set the working directory to /app
WORKDIR /app

# Copy the local package files to the container's workspace.
# Add the source code:
COPY . .

# Build the binary
RUN go build -o main .

# Run the app
CMD ["./main"]
