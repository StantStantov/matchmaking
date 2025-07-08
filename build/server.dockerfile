FROM golang:alpine AS base
WORKDIR /src
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
  go mod download

FROM base AS build
COPY . .
RUN --mount=type=cache,target=/root/.cache/go-build \
  go build -v -o ./tmp/server ./cmd/...

FROM alpine:latest AS runner
WORKDIR /app
COPY --from=build /src/tmp/server ./cmd/server
ENTRYPOINT ["/app/cmd/server"]
