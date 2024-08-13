FROM alpine:latest

# install curl
RUN apk add --update curl && rm -rf /var/cache/apk/*

RUN adduser -D app

# Cd into the api code directory

RUN mkdir -p /go/src/iris
WORKDIR /go/src/iris

# Copy the local package files to the container's workspace.
COPY main conf/ /go/src/iris/

# Chown the application directory to app user
RUN chown -R app:app /go/src/iris/

# Create user's home directory
RUN mkdir -p /home/app
RUN chown app /home/app

# Use the unprivileged user
USER app

HEALTHCHECK --interval=5s --timeout=3s \
CMD curl -fs http://localhost:8090/rest/keepalive || exit 1
ENTRYPOINT ["./main"]

# Document that the service listens on port 8080.
EXPOSE 8090
