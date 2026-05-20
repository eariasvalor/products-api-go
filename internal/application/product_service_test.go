package application

import (
	"errors"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll_Success(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	expected := []*domain.Product{
		{ID: 1, Name: "Silla", Price: 99.99, Stock: 10},
		{ID: 2, Name: "Mesa", Price: 199.99, Stock: 5},
	}

	repo.On("FindAll").Return(expected, nil)

	products, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(products))
	repo.AssertExpectations(t)
}

func TestGetAll_Error(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	repo.On("FindAll").Return(nil, errors.New("error de base de datos"))

	products, err := service.GetAll()

	assert.Error(t, err)
	assert.Nil(t, products)
	repo.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	expected := &domain.Product{ID: 1, Name: "Silla", Price: 99.99, Stock: 10}
	repo.On("FindByID", 1).Return(expected, nil)

	product, err := service.GetByID(1)

	assert.NoError(t, err)
	assert.Equal(t, "Silla", product.Name)
	repo.AssertExpectations(t)
}

func TestGetByID_NotFound(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	repo.On("FindByID", 99).Return(nil, errors.New("no encontrado"))

	product, err := service.GetByID(99)

	assert.Error(t, err)
	assert.Nil(t, product)
	repo.AssertExpectations(t)
}

func TestCreate_Success(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	input := &domain.Product{Name: "Silla", Description: "Silla de oficina", Price: 99.99, Stock: 10}
	expected := &domain.Product{ID: 1, Name: "Silla", Description: "Silla de oficina", Price: 99.99, Stock: 10}

	repo.On("Save", input).Return(expected, nil)

	product, err := service.Create(input)

	assert.NoError(t, err)
	assert.Equal(t, 1, product.ID)
	repo.AssertExpectations(t)
}

func TestCreate_InvalidProduct(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	input := &domain.Product{Name: "", Price: -10, Stock: -1}

	product, err := service.Create(input)

	assert.Error(t, err)
	assert.Nil(t, product)
	repo.AssertNotCalled(t, "Save")
}

func TestDelete_Success(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	repo.On("Delete", 1).Return(nil)

	err := service.Delete(1)

	assert.NoError(t, err)
	repo.AssertExpectations(t)
}

func TestDelete_NotFound(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)
	service := NewProductService(repo)

	repo.On("Delete", 99).Return(errors.New("producto no encontrado"))

	err := service.Delete(99)

	assert.Error(t, err)
	repo.AssertExpectations(t)
}
