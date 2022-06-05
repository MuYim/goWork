package services

import (
	"context"
)

type ProdService struct {
	UnimplementedProdServiceServer
}

func (c *ProdService) GetProdStock(context context.Context, request *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}
