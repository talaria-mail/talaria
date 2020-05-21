FROM node:12 as frontend
WORKDIR /scratch
COPY frontend .
RUN yarn install
RUN yarn build

FROM golang as backend
WORKDIR /scratch
RUN go get github.com/go-bindata/go-bindata/...
COPY . .
COPY --from=frontend /scratch/dist frontend/dist
RUN go generate cmd/talaria/*
RUN CGO_ENABLED=0 go build -o talaria cmd/talaria/*

FROM scratch
ENV PATH /bin
COPY --from=backend /scratch/talaria /bin/talaria
EXPOSE 8080
ENTRYPOINT ["talaria", "server"]
