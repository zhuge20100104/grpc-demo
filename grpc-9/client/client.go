package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/grpc-demo/grpc-9/client/helper"

	"github.com/grpc-demo/grpc-9/client/services"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()

	orderClient := services.NewOrderServiceClient(conn)
	orderTime := timestamp.Timestamp{Seconds: time.Now().Unix()}
	orderRes, err := orderClient.NewOrder(context.Background(),
		&services.OrderRequest{
			OrderMain: &services.OrderMain{
				OrderId:    100,
				OrderNo:    "酱油1001",
				UserId:     200,
				OrderMoney: 9.8,
				OrderTime:  &orderTime},
		})

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	fmt.Printf("订单状态 %v, 订单消息: %v\n", orderRes.Status, orderRes.Message)
}
