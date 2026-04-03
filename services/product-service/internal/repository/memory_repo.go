package repository

import (
	"errors"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/domain"
	"sync"
)

// MemoryProductRepo hafızada çalışan bir repository (Testler için ideal)
type MemoryProductRepo struct {
	mu       sync.RWMutex
	products map[string]*domain.Product
}

func NewMemoryProductRepo() *MemoryProductRepo {
	return &MemoryProductRepo{
		products: make(map[string]*domain.Product),
	}
}

func (r *MemoryProductRepo) Create(p *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.products[p.ID] = p
	return nil
}

func (r *MemoryProductRepo) GetByID(id string) (*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.products[id]
	if !ok {
		return nil, errors.New("urun bulunamadi")
	}
	return p, nil
}

func (r *MemoryProductRepo) List() ([]*domain.Product, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*domain.Product, 0, len(r.products))
	for _, p := range r.products {
		list = append(list, p)
	}
	return list, nil
}
