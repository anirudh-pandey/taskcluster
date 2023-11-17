# Simple generic worker

FROM golang:1.21.4-bookworm AS build

WORKDIR /app

# build depends on the .git
COPY . .

ENV CGO_ENABLED=0
RUN cd tools/livelog && go build -o /livelog && cd ..
RUN cd tools/worker-runner && go build -o /start-worker ./cmd/start-worker && cd ..
RUN cd tools/taskcluster-proxy && go build -o /taskcluster-proxy && cd ..
RUN cd clients/client-shell && go build -o /taskcluster && cd ../..
RUN cd workers/generic-worker && \
  ./build.sh && \
  mv generic-worker-multiuser-* /generic-worker-multiuser && \
  mv generic-worker-simple-* /generic-worker


FROM ubuntu:jammy

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install -y ca-certificates curl gzip

COPY --from=build /livelog /taskcluster-proxy /start-worker /taskcluster /generic-worker* /usr/local/bin/
RUN ls -la /usr/local/bin

RUN mkdir -p /etc/generic-worker
RUN mkdir -p /var/local/generic-worker

# autogenerated ed25519 key is only good for local and testing
RUN generic-worker new-ed25519-keypair --file /etc/generic-worker/ed25519_key

# Write out the DockerFlow-compatible version.json file
ARG DOCKER_FLOW_VERSION
RUN if [ -n "${DOCKER_FLOW_VERSION}" ]; then \
    echo "${DOCKER_FLOW_VERSION}" > /version.json; \
else \
    echo \{\"version\": \"58.0.0\", \"commit\": \"local\", \"source\": \"https://github.com/taskcluster/taskcluster\", \"build\": \"NONE\"\} > /version.json; \
fi

VOLUME /etc/generic-worker/config.json
VOLUME /var/local/generic-worker

COPY workers/entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
WORKDIR /var/local/generic-worker

ENTRYPOINT [ "/entrypoint.sh" ]
