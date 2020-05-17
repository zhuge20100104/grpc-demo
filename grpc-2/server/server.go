package main

import (
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	"github.com/grpc-demo/grpc-2/server/services"
	"google.golang.org/grpc"
)

func main() {

	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	if err != nil {
		log.Fatalf("加载服务端证书和Key失败, err: %v\n", err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}
	rpcServer.Serve(listen)
}
