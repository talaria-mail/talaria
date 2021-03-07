FROM golang:1.16 as build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o talaria cmd/talaria/*

FROM scratch as prod
WORKDIR /bin
COPY --from=build /build/talaria .
CMD ["/bin/talaria", "serve"]
