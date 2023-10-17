FROM golang:1.21-alpine AS builder
ARG GIT_PRIVATE_REPO_URL
ENV GO111MODULE=on
ENV GOPRIVATE=${GIT_PRIVATE_REPO_URL}
RUN apk update && apk add git make bash
WORKDIR /app
COPY go.* .
RUN --mount=type=secret,id=netrc,dst=/root/.netrc go mod download
COPY . .
RUN make build-payment

FROM alpine:3.17.3 AS runner
WORKDIR /app
COPY --from=builder ./app/.bin .
RUN addgroup -S payment && \
    adduser -S payment -G payment && \
    chown -R payment:payment . && \
    chmod -R 500 .
USER payment
CMD ["./pls-payment"]