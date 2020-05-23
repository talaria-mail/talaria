# Talaria | Web client

This directory hosts the web client for Talaria. The web client is based on a
[Vue](vuejs.org) single page application. Communication with the backend API is
done via GRPC.

To develop, generate the GRPC server stubs, install dependencies and spin up
a development server:

```shell
$ yarn generate
$ yarn install
$ yarn serve
```

To generate a production build, run

```shell
$ yarn build
```
