package services

import (
	context "context"
	"fmt"
)

type OrderService struct{}

func (*OrderService) NewOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error) {
	fmt.Println(orderRequest.OrderMain)

	return &OrderResponse{
		Status:  "OK",
		Message: "创建主订单成功",
	}, nil
}
