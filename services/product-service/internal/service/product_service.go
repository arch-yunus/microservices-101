package service

import (
	"errors"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/domain"
	"time"
)

// ProductService iş mantığını yöneten ana yapı (Application Layer)
type ProductService struct {
	repo domain.ProductRepository
}

// NewProductService yeni bir servis oluşturur
func NewProductService(repo domain.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

// CreateNewProduct iş kurallarını kontrol eder ve ürünü oluşturur
func (s *ProductService) CreateNewProduct(name string, price float64) (*domain.Product, error) {
	if name == "" {
		return nil, errors.New("urun ismi bos olamaz")
	}
	if price <= 0 {
		return nil, errors.New("fiyat sifirdan buyuk olmalidir")
	}

	product := &domain.Product{
		ID:        "fake-uuid", // Gerçek sistemde UUID gen edilir
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(product)
	return product, err
}
