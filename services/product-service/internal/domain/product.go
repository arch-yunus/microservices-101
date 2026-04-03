package domain

import "time"

// Product temsil eden ana domain nesnesi
type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

// ProductRepository veriye erişim için gerekli arayüz (Interface)
// Bu sayede veritabanı değişimlerinden etkilenmeyiz (Dependency Inversion)
type ProductRepository interface {
	GetByID(id string) (*Product, error)
	Create(product *Product) error
	List() ([]*Product, error)
}
