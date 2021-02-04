<p align="center">
  <img src="./docs/assets/talaria.svg" alt="Talaria" width="250px"/>
</p>

# Talaria

[![Build Status][drone-badge]][drone-link]
[![Container Layers][micro-badger-layers]][micro-badger]
[![Container Image][micro-badger-version]][micro-badger]

[micro-badger]: https://microbadger.com/images/nsmith5/talaria
[micro-badger-layers]: https://images.microbadger.com/badges/image/nsmith5/talaria.svg
[micro-badger-version]: https://images.microbadger.com/badges/version/nsmith5/talaria.svg

Talaria is an effort to create an email server that goes out of its way to make
it easy for you to host your own email.

**Goals**

- Low resource usage (1 vCPU, 500MiB Ram should be able to comfortably run Talaria)
- Easy configuration (Don't rely on docs to get users to set up tricky DNS, ask for AWS creds (or equivalent) and go set it up for them)
- Target only modern protocols (implicit TLS on submission, implicit TLS on IMAP, no support of POP at all etc)
- Stay away from email black lists with rigorous compliance to DKIM, SPF and other identity protocols (without making the user think about this stuff!)

**Non-Goals**

- Exhaustive compliance with all protocols
- High scalability or high availability

## Build & Run

To run the lastest container image:

```
$ docker run -p 8080:8080 -p 8081:8081 nsmith5/talaria
```

To compile from source:

```shell
$ git clone https://github.com/nsmith5/talaria
$ pushd talaria
$ pushd frontend
$ yarn install
$ yarn build
$ popd
$ go get github.com/go-bindata/go-bindata/...
$ go generate cmd/talaria/*    # Use go-bindata to bundle frontend into go binary
$ go build -o talaria cmd/talaria/*
$ ./talaria server
```

[drone-badge]: https://cloud.drone.io/api/badges/nsmith5/talaria/status.svg
[drone-link]: https://cloud.drone.io/nsmith5/talaria
