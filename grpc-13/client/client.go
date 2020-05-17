package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/grpc-demo/grpc-13/client/helper"

	"github.com/grpc-demo/grpc-13/client/services"

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

	stream, err := userClient.GetUserScoreByServerStream(context.Background(),
		&services.UserScoreRequest{Users: users},
	)

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	for {
		userRes, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("读取服务端流失败 err: %v\n", err.Error())
		}
		fmt.Println(userRes.Users)
	}
}
