package output

import "proyectoGo/internal/domain"

type ProductRepository interface {
	FindAll() ([]*domain.Product, error)
	FindByID(id int) (*domain.Product, error)
	Save(product *domain.Product) (*domain.Product, error)
	Update(id int, product *domain.Product) (*domain.Product, error)
	Delete(id int) error
}
