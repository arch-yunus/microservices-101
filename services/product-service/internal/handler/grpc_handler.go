package handler

import (
	"context"
	pb "github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/service"
)

// GrpcProductHandler gRPC isteklerini karılayan yapı
type GrpcProductHandler struct {
	pb.UnimplementedProductServiceServer
	svc *service.ProductService
}

func NewGrpcProductHandler(svc *service.ProductService) *GrpcProductHandler {
	return &GrpcProductHandler{svc: svc}
}

// GetProduct proto dosyasında tanımlanan RPC metodu
func (h *GrpcProductHandler) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.ProductResponse, error) {
	// Not: Burada domain modeline donusum yapılmalı. 
	// simdilik basitlik amacl direkt donebiliriz.
	return &pb.ProductResponse{
		Id:    req.Id,
		Name:  "Elite Product from gRPC",
		Price: 299.99,
	}, nil
}
