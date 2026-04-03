package service

import (
	"context"
	"fmt"
	"github.com/arch-yunus/microservices-101/proto/product"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/domain"
	"time"
)

// OrderService sipari i mantn yöneten yapı
type OrderService struct {
	productClient product.ProductServiceClient
}

func NewOrderService(pc product.ProductServiceClient) *OrderService {
	return &OrderService{productClient: pc}
}

// PlaceOrder yeni bir sipari olusturur
func (s *OrderService) PlaceOrder(ctx context.Context, productID string, quantity int) (*domain.Order, error) {
	// 1. gRPC ile urun bilgilerini sorda
	p, err := s.productClient.GetProduct(ctx, &product.GetProductRequest{Id: productID})
	if err != nil {
		return nil, fmt.Errorf("urun bilgisi alinamadi: %v", err)
	}

	// 2. Sipari olustur
	order := &domain.Order{
		ID:        "order-123",
		ProductID: p.Id,
		Quantity:  quantity,
		Total:     p.Price * float64(quantity),
		CreatedAt: time.Now(),
	}

	fmt.Printf("?? Sipari Olusturuldu: %s iin %d adet (Toplam: %.2f)\n", p.Name, quantity, order.Total)
	return order, nil
}
