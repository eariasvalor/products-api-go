package application

import (
	"errors"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/ports/input"
	"proyectoGo/internal/ports/output"
)

var _ input.ProductService = (*ProductService)(nil)

type ProductService struct {
	repo output.ProductRepository
}

func NewProductService(repo output.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll() ([]*domain.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetByID(id int) (*domain.Product, error) {
	product, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("producto no encontrado")
	}
	return product, nil
}

func (s *ProductService) Create(product *domain.Product) (*domain.Product, error) {
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return s.repo.Save(product)
}

func (s *ProductService) Update(id int, product *domain.Product) (*domain.Product, error) {
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return s.repo.Update(id, product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
