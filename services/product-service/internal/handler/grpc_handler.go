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
	product, err := h.svc.GetProduct(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.String(),
	}, nil
}

// CreateProduct proto dosyasında tanımlanan RPC metodu
func (h *GrpcProductHandler) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.ProductResponse, error) {
	product, err := h.svc.CreateNewProduct(req.Name, req.Description, req.Price)
	if err != nil {
		return nil, err
	}

	return &pb.ProductResponse{
		Id:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt.String(),
	}, nil
}
