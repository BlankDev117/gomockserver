FROM golang:1.15.6-alpine3.12 as buildStage
WORKDIR /usr/local/go/src/github.com/BlankDev117/gomockserver/src
COPY src .
RUN go build -o "./bin/gomockserver"

FROM alpine:3.12 as deployStage
COPY --from=buildStage /usr/local/go/src/github.com/BlankDev117/gomockserver/src/bin/gomockserver /gomockserver
ENV PORT=8080
ENTRYPOINT ["/gomockserver"]