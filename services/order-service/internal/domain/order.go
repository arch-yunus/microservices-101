package domain

import "time"

type Order struct {
	ID        string    `json:"id" param:"id"`
	ProductID string    `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderRepository interface {
	Create(order *Order) error
	GetByID(id string) (*Order, error)
}
