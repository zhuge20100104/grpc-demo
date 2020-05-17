package main

import (
	"context"
	"fmt"
	"log"

	"github.com/grpc-demo/grpc-12/client/helper"

	"github.com/grpc-demo/grpc-12/client/services"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8081", grpc.WithTransportCredentials(helper.GetClientCredentials()))
	if err != nil {
		log.Fatalf("连接GRPC服务端失败 %v\n", err)
	}

	defer conn.Close()

	userClient := services.NewUserServiceClient(conn)

	users := make([]*services.UserInfo, 0)
	var i int32 = 0
	for i = 0; i < 6; i++ {
		user := &services.UserInfo{UserId: i + 1}
		users = append(users, user)
	}

	userRes, err := userClient.GetUserScore(context.Background(),
		&services.UserScoreRequest{Users: users},
	)

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	fmt.Println(userRes.Users)
}
