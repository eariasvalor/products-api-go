package mocks

import (
	"proyectoGo/internal/domain"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) FindAll() ([]*domain.Product, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*domain.Product), args.Error(1)
}

func (m *ProductRepositoryMock) FindByID(id int) (*domain.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Save(product *domain.Product) (*domain.Product, error) {
	args := m.Called(product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Update(id int, product *domain.Product) (*domain.Product, error) {
	args := m.Called(id, product)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.Product), args.Error(1)
}

func (m *ProductRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
