package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/credentials"

	"github.com/grpc-demo/grpc-2/client/services"

	"google.golang.org/grpc"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("keys/server.crt", "sensetime.com")
	if err != nil {
		log.Fatalf("加载客户端证书失败, err: %v\n", err)
	}

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()
	prodClient := services.NewProductServiceClient(conn)
	prodRes, err := prodClient.GetProductStock(context.Background(),
		&services.ProdRequest{ProdId: 12})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	fmt.Println(prodRes.ProdStock)
}
