# syntax=docker/dockerfile:1.2
#
# firearweave builder
FROM golang:buster as firearweave-builder
COPY . firehose-arweave
RUN --mount=type=cache,target=/var/cache/apk \
    --mount=type=cache,target=/go/pkg \
    cd firehose-arweave \
    && go install -v -ldflags "-X main.Version=$version -X main.Commit=`git rev-list -1 HEAD`" \
    ./cmd/firearweave

# thegarii builder
FROM rust:buster as thegarii-builder
ENV CARGO_NET_GIT_FETCH_WITH_CLI=true
RUN --mount=type=cache,target=/var/cache/apk \
    --mount=type=cache,target=/home/rust/.cargo \
    rustup component add rustfmt \
    && cargo install thegarii

# firearweave
FROM debian:stable-slim as firearweave-release
COPY --from=firearweave-builder /go/bin/firearweave /usr/bin/firearweave
COPY --from=thegarii-builder /usr/local/cargo/bin/thegarii /usr/bin/thegarii
COPY ./devel/standard/standard.yaml config.yaml
RUN apt-get update \
    && apt-get install ca-certificates -y \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["firearweave", "-c", "config.yaml", "start"]

