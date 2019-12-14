FROM golang as builder
COPY . /go/src/proj
RUN cd /go/src/proj && GOCACHE=/opt CGO_ENABLED=0 go build -o talaria ./cmd/talaria/*

FROM scratch
COPY --from=builder /go/src/proj/talaria /
ENTRYPOINT ["/talaria"]
