# syntax=docker/dockerfile:1.2
#
# firearweave builder
FROM golang:alpine as firearweave-builder
COPY . firehose-arweave
RUN --mount=type=cache,target=/var/cache/apk \
    --mount=type=cache,target=/go/pkg \
    apk add git \
    && cd firehose-arweave \
    && sed -i 's/\/home\/clearloop/\/go/g' go.mod \
    && go install -v -ldflags "-X main.Version=$version -X main.Commit=`git rev-list -1 HEAD`" \
    ./cmd/firearweave

# thegarii builder
FROM rust:alpine as thegarii-builder
ENV CARGO_NET_GIT_FETCH_WITH_CLI=true
RUN --mount=type=cache,target=/var/cache/apk \
    --mount=type=cache,target=/home/rust/.cargo \
    apk add git musl-dev openssl-dev protoc \
    && rustup component add rustfmt \
    && cargo install thegarii

# firearweave
FROM alpine as firearweave-release
COPY --from=firearweave-builder /go/bin/firearweave /usr/bin/firearweave
COPY --from=thegarii-builder /usr/local/cargo/bin/thegarii /usr/bin/thegarii
COPY ./devel/standard/standard.yaml config.yaml
CMD ["firearweave", "-c", "config.yaml", "start"]
