package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-demo/grpc-6/client/helper"

	"github.com/grpc-demo/grpc-6/client/services"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()

	prodClient := services.NewProductServiceClient(conn)
	prod, err := prodClient.GetProductStock(context.Background(),
		&services.ProdRequest{ProdId: 20, ProdArea: services.ProdAreas_B})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	fmt.Println(prod.ProdStock)
}
