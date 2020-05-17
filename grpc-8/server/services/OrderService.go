package services

import (
	context "context"
	"fmt"
)

type OrderService struct{}

func (*OrderService) NewOrder(ctx context.Context, orderMain *OrderMain) (*OrderResponse, error) {
	fmt.Println(orderMain)
	return &OrderResponse{
		Status:  "OK",
		Message: "创建主订单成功",
	}, nil
}
