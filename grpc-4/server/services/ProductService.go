package services

import context "context"

type ProdService struct{}

func (*ProdService) GetProductStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{
		ProdStock: 300,
	}, nil
}
