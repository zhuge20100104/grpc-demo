package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhuge20100104/grpc-demo/grpc-7/client/helper"

	"github.com/zhuge20100104/grpc-demo/grpc-7/client/services"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()

	prodClient := services.NewProductServiceClient(conn)
	prod, err := prodClient.GetProductInfo(context.Background(),
		&services.ProdRequest{ProdId: 20})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	fmt.Printf("ID: %v, Name: %v, Price: %v\n", prod.ProdId,
		prod.ProdName, prod.ProdPrice)
}
