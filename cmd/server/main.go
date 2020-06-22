package main

import (
	"google.golang.org/grpc"
	"net"
	"simple-adder-grpc-golang/pkg/add"
	"simple-adder-grpc-golang/pkg/api"
)

func main() {
	server := grpc.NewServer()
	addSrv := &add.AddService{}

	api.RegisterAddServer(server, addSrv)

	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	server.Serve(l)
}