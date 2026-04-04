package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/arch-yunus/microservices-101/services/order-service/internal/domain"
)

type PostgresOrderRepository struct {
	db *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) *PostgresOrderRepository {
	return &PostgresOrderRepository{db: db}
}

func (r *PostgresOrderRepository) Create(o *domain.Order) error {
	query := `INSERT INTO orders (id, product_id, quantity, total, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, o.ID, o.ProductID, o.Quantity, o.Total, o.CreatedAt)
	return err
}

func (r *PostgresOrderRepository) GetByID(id string) (*domain.Order, error) {
	query := `SELECT id, product_id, quantity, total, created_at FROM orders WHERE id = $1`
	var o domain.Order
	err := r.db.QueryRow(query, id).Scan(&o.ID, &o.ProductID, &o.Quantity, &o.Total, &o.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("sipariş bulunamadı: %s", id)
		}
		return nil, err
	}
	return &o, nil
}
