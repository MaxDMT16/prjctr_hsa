# syntax=docker/dockerfile:1

FROM golang:1.21

ENV API_PORT=80
ENV REDIS_ADDRESS=localhost:6379
ENV BEANSTALK_ADDRESS=localhost:11300

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./
COPY cmd/*.go cmd/

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /queues-app ./cmd/main.go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE $API_PORT

# Run
CMD ["/queues-app"]