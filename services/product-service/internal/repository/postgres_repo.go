package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/arch-yunus/microservices-101/services/product-service/internal/domain"
)

// PostgresProductRepository PostgreSQL veritabanı iletişimi için yapı
type PostgresProductRepository struct {
	db *sql.DB
}

// NewPostgresProductRepository yeni bir veri tabanı deposu oluşturur
func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) GetByID(id string) (*domain.Product, error) {
	query := `SELECT id, name, description, price, created_at FROM products WHERE id = $1`
	var p domain.Product
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("ürün bulunamadı: %s", id)
		}
		return nil, err
	}
	return &p, nil
}

func (r *PostgresProductRepository) Create(p *domain.Product) error {
	query := `INSERT INTO products (id, name, description, price, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, p.ID, p.Name, p.Description, p.Price, p.CreatedAt)
	return err
}

func (r *PostgresProductRepository) List() ([]*domain.Product, error) {
	query := `SELECT id, name, description, price, created_at FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		var p domain.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.CreatedAt); err != nil {
			return nil, err
		}
		products = append(products, &p)
	}
	return products, nil
}
