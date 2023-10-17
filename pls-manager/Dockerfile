FROM golang:1.21-alpine AS builder
ARG GIT_PRIVATE_REPO_URL
ENV GO111MODULE=on
ENV GOPRIVATE=${GIT_PRIVATE_REPO_URL}
RUN apk update && apk add git make bash
WORKDIR /app
COPY go.* .
RUN --mount=type=secret,id=netrc,dst=/root/.netrc go mod download
COPY . .
RUN make build-manager

FROM alpine:3.17.3 AS runner
WORKDIR /app
COPY --from=builder ./app/.bin .
RUN addgroup -S manager && \
    adduser -S manager -G manager && \
    chown -R manager:manager . && \
    chmod -R 500 .
USER manager
CMD ["./pls-manager"]