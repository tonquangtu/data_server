package main

import (
	"fmt"
	"github.com/tonquangtu/data_server/internal"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	container, err := internal.NewContainer()
	if err != nil {
		fmt.Println("Cannot start server", err)
		return
	}

	// create a custom RPC server
	server := rpc.NewServer()

	// register `mit` object with `server`
	server.Register(container.UserService)

	// create a TCP listener at address : 127.0.0.1:9002
	// https://golang.org/pkg/net/#Listener
	listener, _ := net.Listen("tcp", "127.0.0.1:6969")

	for {
		// get connection from the listener when client connects
		conn, _ := listener.Accept() // Accept blocks until next connection is received

		// serve connection in a separate goroutine using JSON codec
		go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
