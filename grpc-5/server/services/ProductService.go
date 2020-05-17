package services

import context "context"

type ProdService struct{}

func (*ProdService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 300,
	}, nil
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
