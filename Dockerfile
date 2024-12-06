# syntax=docker/dockerfile:1

FROM golang:1.22.2-alpine AS Builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


WORKDIR /build

#RUN apk update && apk upgrade && \
#    apk --no-cache add ca-certificates && \
#    apk --no-cache add tzdata

COPY . .
RUN go mod download
RUN sleep 10
RUN go build -a -o start_point cmd/main.go

FROM alpine:3.20

#RUN apk update && \
#    apk --no-cache add curl && \
#    apk add --no-cache tzdata

WORKDIR /app
COPY --from=Builder /build/start_point .

EXPOSE 8080
ENTRYPOINT [ "./start_point" ]