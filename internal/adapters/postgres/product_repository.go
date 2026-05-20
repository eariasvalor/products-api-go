package postgres

import (
	"database/sql"
	"errors"
	"log"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/ports/output"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var _ output.ProductRepository = (*ProductRepository)(nil)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) FindAll() ([]*domain.Product, error) {
	rows, err := r.db.Query("SELECT id, name, description, price, stock FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*domain.Product
	for rows.Next() {
		p := &domain.Product{}
		err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id int) (*domain.Product, error) {
	p := &domain.Product{}
	err := r.db.QueryRow(
		"SELECT id, name, description, price, stock FROM products WHERE id = $1", id,
	).Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Stock)

	if err == sql.ErrNoRows {
		return nil, errors.New("producto no encontrado")
	}
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *ProductRepository) Save(product *domain.Product) (*domain.Product, error) {
	var dbName string
	r.db.QueryRow("SELECT current_database()").Scan(&dbName)
	log.Println("Base de datos activa:", dbName)

	err := r.db.QueryRow(
		"INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) RETURNING id",
		product.Name, product.Description, product.Price, product.Stock,
	).Scan(&product.ID)

	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r *ProductRepository) Update(id int, product *domain.Product) (*domain.Product, error) {
	result, err := r.db.Exec(
		"UPDATE products SET name=$1, description=$2, price=$3, stock=$4 WHERE id=$5",
		product.Name, product.Description, product.Price, product.Stock, id,
	)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, errors.New("producto no encontrado")
	}

	product.ID = id
	return product, nil
}

func (r *ProductRepository) Delete(id int) error {
	result, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("producto no encontrado")
	}
	return nil
}
