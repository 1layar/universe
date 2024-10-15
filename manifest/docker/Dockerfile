FROM golang:1.23.0-alpine AS base
WORKDIR /app

# builder
FROM base AS builder
ENV GOOS=linux
ENV GOARCH=amd64

# build-args
ARG VERSION
ARG SERVICE_NAME

RUN apk --no-cache add bash git openssh

# modules: utilize build cache
COPY . .

# Sync the dependencies
RUN go work sync

# inject versioning information & build the binary
# Build the application
RUN export BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ") && \
    go build -o app \
    -ldflags "-X github.com/1layar/universe/internal/${SERVICE_NAME}/app/appbundle.Version=${VERSION} \
              -X github.com/1layar/universe/internal/${SERVICE_NAME}/app/appbundle.BuildTimeString=${BUILD_TIME}" \
    ./internal/${SERVICE_NAME}

# runner
FROM base AS runner
RUN apk add --no-cache libc6-compat tini
# Tini is now available at /sbin/tini

COPY --from=builder /app/app /app/app
EXPOSE 8080

ENTRYPOINT ["/sbin/tini", "--"]
CMD [ "/app/app", "start" ]
