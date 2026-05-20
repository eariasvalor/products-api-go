package application

import (
	"errors"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStockMonitor_DetectsLowStockProducts(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)

	products := []*domain.Product{
		{ID: 1, Name: "Silla", Price: 99.99, Stock: 2},
		{ID: 2, Name: "Mesa", Price: 199.99, Stock: 20},
	}

	repo.On("FindAll").Return(products, nil)

	monitor := NewStockMonitor(repo, 5)
	lowStock := monitor.CheckLowStock()

	assert.Len(t, lowStock, 1)
	assert.Equal(t, "Silla", lowStock[0].Name)
	repo.AssertExpectations(t)
}

func TestStockMonitor_NoLowStockProducts(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)

	products := []*domain.Product{
		{ID: 1, Name: "Silla", Price: 99.99, Stock: 10},
		{ID: 2, Name: "Mesa", Price: 199.99, Stock: 20},
	}

	repo.On("FindAll").Return(products, nil)

	monitor := NewStockMonitor(repo, 5)
	lowStock := monitor.CheckLowStock()

	assert.Empty(t, lowStock)
	repo.AssertExpectations(t)
}

func TestStockMonitor_RepoError(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)

	repo.On("FindAll").Return(nil, errors.New("error de base de datos"))

	monitor := NewStockMonitor(repo, 5)
	lowStock := monitor.CheckLowStock()

	assert.Nil(t, lowStock)
	repo.AssertExpectations(t)
}

func TestStockMonitor_StartsAndStops(t *testing.T) {
	repo := new(mocks.ProductRepositoryMock)

	products := []*domain.Product{
		{ID: 1, Name: "Silla", Price: 99.99, Stock: 2},
	}

	repo.On("FindAll").Return(products, nil)

	monitor := NewStockMonitor(repo, 5)
	stop := monitor.Start(100 * time.Millisecond)

	time.Sleep(350 * time.Millisecond)
	stop()

	repo.AssertNumberOfCalls(t, "FindAll", 3)
}
