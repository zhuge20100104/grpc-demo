package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/zhuge20100104/grpc-demo/grpc-7/server/helper"
	"github.com/zhuge20100104/grpc-demo/grpc-7/server/services"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {

	rpcServer := grpc.NewServer(grpc.Creds(helper.GetServerCredentials()))
	services.RegisterProductServiceServer(rpcServer, new(services.ProdService))
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("启动网络监听失败 %v\n", err)
	}

	ch := make(chan int, 2)
	go func() {
		rpcServer.Serve(listen)
		ch <- 1
	}()

	// 使用Gateway启动HTTP Server
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCredentials())}
	err = services.RegisterProductServiceHandlerFromEndpoint(context.Background(),
		gwmux, "localhost:8081", opt)

	if err != nil {
		log.Fatalf("从GRPC-GateWay连接GRPC失败, err: %v\n", err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	go func() {
		httpServer.ListenAndServe()
		ch <- 1
	}()

	for i := 0; i < 2; i++ {
		<-ch
	}

}
