FROM golang:alpine AS build-env
ADD . /src
RUN apk add --update --no-cache git && \
    go get github.com/prometheus/client_golang/prometheus/promhttp && \
    cd /src && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o main main.go

FROM scratch
ARG BUILD_DATE
ARG VCS_REF
LABEL org.label-schema.schema-version=1.0 \
      org.label-schema.build-date=$BUILD_DATE \
      org.label-schema.vcs-ref=$VCS_REF \
      org.label-schema.vcs-url="https://github.com/nbarnum/docker-memory-leak"
COPY --from=build-env /src/main /
EXPOSE 8081
ENTRYPOINT ["/main"]
