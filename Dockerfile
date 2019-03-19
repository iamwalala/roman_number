# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.12
ARG GO_VERSION=1.12

# First stage: build the executable.
FROM golang:${GO_VERSION}

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/iamwalala/

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Install the package
RUN go install -v  github.com/iamwalala/romanserver

ENV PATH=$PATH:$GOPATH/bin

# Run the executable
CMD ["romanserver"]