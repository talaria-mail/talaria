package proto

//go:generate protoc --go_out=plugins=grpc:. auth.proto
//go:generate protoc --go_out=plugins=grpc:. users.proto
