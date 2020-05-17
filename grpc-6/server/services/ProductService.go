package services

import context "context"

type ProdService struct{}

func (*ProdService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.ProdArea {
	case ProdAreas_A:
		stock = 10
	case ProdAreas_B:
		stock = 20
	case ProdAreas_C:
		stock = 30
	default:
		stock = 0
	}

	return &ProdResponse{ProdStock: stock}, nil
}

func (*ProdService) GetProductStocks(ctx context.Context, req *QuerySize) (*ProdResponseList, error) {
	size := req.GetSize()
	res := make([]*ProdResponse, 0)
	var i int32 = 0
	for i = 0; i < size; i++ {
		res = append(res, &ProdResponse{ProdStock: i + 1})
	}
	return &ProdResponseList{Prodres: res}, nil
}
