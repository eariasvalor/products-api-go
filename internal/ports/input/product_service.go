package input

import "proyectoGo/internal/domain"

type ProductService interface {
	GetAll() ([]*domain.Product, error)
	GetByID(id int) (*domain.Product, error)
	Create(product *domain.Product) (*domain.Product, error)
	Update(id int, product *domain.Product) (*domain.Product, error)
	Delete(id int) error
}
