# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/Shelnutt2/galera_http_check

# Change working directory to package
WORKDIR $GOPATH/src/github.com/Shelnutt2/galera_http_check

# Install dependencies
RUN go get -v ./...

# Build the galera_http_check command inside the container.
RUN go install github.com/Shelnutt2/galera_http_check

# Set env for galera_http_check port
ENV HTTP_PORT 80

#Set env for galera host
ENV GALERA_HOST localhost

# Run the galera_http_check command by default when the container starts.
CMD /go/bin/galera_http_check --port $HTTP_PORT --host $GALERA_HOST

# Document that the service listens on port 80.
EXPOSE 80
