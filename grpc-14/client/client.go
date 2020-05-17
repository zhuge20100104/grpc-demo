package main

import (
	"context"
	"fmt"
	"log"

	"github.com/zhuge20100104/grpc-demo/grpc-14/client/helper"

	"github.com/zhuge20100104/grpc-demo/grpc-14/client/services"

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

	stream, err := userClient.GetUserScoreByClientStream(context.Background())

	if err != nil {
		log.Fatalf("请求GRPC服务端失败 %v\n", err)
	}

	for i := 0; i < 3; i++ {
		req := new(services.UserScoreRequest)
		req.Users = make([]*services.UserInfo, 0)
		var j int32
		for j = 1; j <= 5; j++ {
			req.Users = append(req.Users, &services.UserInfo{UserId: j})
		}
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("接收服务端请求失败 %v\n", err)
	}

	for _, user := range res.Users {
		fmt.Println(user)
	}

}
