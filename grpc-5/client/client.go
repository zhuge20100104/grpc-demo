package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-demo/grpc-5/client/helper"

	"github.com/grpc-demo/grpc-5/client/services"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()

	prodClient := services.NewProductServiceClient(conn)
	prods, err := prodClient.GetProductStocks(context.Background(),
		&services.QuerySize{Size: 20})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}
	for _, res := range prods.Prodres {
		fmt.Printf("商品库存为: %v\n", res.ProdStock)
	}
}
