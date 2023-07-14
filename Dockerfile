FROM golang:alpine as builder
WORKDIR $GOPATH/src/
COPY . .
RUN go build -o /go/bin/probe server.go

FROM scratch
COPY --from=builder /go/bin/probe /go/bin/probe
EXPOSE 5341
ENTRYPOINT ["/go/bin/probe"]