FROM golang:1.16.7-alpine3.14 AS golang

ARG LOC=/builds/go/src/github.com/betorvs/playbypost-dnd/
RUN apk add --no-cache git
RUN mkdir -p $LOC
ENV GOPATH /go
COPY . $LOC
ENV CGO_ENABLED 0
RUN cd $LOC && TESTRUN=true go test ./... && go build

FROM alpine:3.14
ARG LOC=/builds/go/src/github.com/betorvs/playbypost-dnd
WORKDIR /
VOLUME /tmp
RUN apk add --no-cache ca-certificates
RUN update-ca-certificates
RUN mkdir -p /app
RUN addgroup -g 1000 -S app && \
    adduser -u 1000 -G app -S -D -h /app app && \
    chmod 755 /app
COPY --from=golang $LOC/playbypost-dnd /app

EXPOSE 8080
RUN chmod +x /app/playbypost-dnd
WORKDIR /app    
USER app
CMD ["/app/playbypost-dnd"]
