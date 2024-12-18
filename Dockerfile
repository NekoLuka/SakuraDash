FROM golang:1.23.3-alpine AS build-stage

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go index.html ./
COPY static ./static
RUN CGO_ENABLED=0 GOOS=linux go build -o /sakuradash

FROM alpine:3.21.0 AS execution-stage

WORKDIR /

COPY --from=build-stage /sakuradash /sakuradash

EXPOSE 8080

ENTRYPOINT ["/sakuradash"]