package services

import (
	context "context"
	"fmt"
)

type OrderService struct{}

func (*OrderService) NewOrder(ctx context.Context, orderRequest *OrderRequest) (*OrderResponse, error) {

	// 加上MainOrder的Validate插件
	err := orderRequest.OrderMain.Validate()
	if err != nil {
		return &OrderResponse{
			Status:  "Error",
			Message: fmt.Sprintf("创建主订单失败: %v\n", err.Error()),
		}, nil
	}

	fmt.Println(orderRequest.OrderMain)

	return &OrderResponse{
		Status:  "OK",
		Message: "创建主订单成功",
	}, nil
}
