package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct_Success(t *testing.T) {
	product, err := NewProduct("Silla", "Silla de oficina", 99.99, 10)

	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "Silla", product.Name)
	assert.Equal(t, 99.99, product.Price)
	assert.Equal(t, 10, product.Stock)
}

func TestNewProduct_EmptyName(t *testing.T) {
	product, err := NewProduct("", "Silla de oficina", 99.99, 10)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "el nombre es obligatorio", err.Error())
}

func TestNewProduct_InvalidPrice(t *testing.T) {
	product, err := NewProduct("Silla", "Silla de oficina", -10, 10)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "el precio debe ser mayor a 0", err.Error())
}

func TestNewProduct_NegativeStock(t *testing.T) {
	product, err := NewProduct("Silla", "Silla de oficina", 99.99, -1)

	assert.Error(t, err)
	assert.Nil(t, product)
	assert.Equal(t, "el stock no puede ser negativo", err.Error())
}

func TestValidate_Success(t *testing.T) {
	product := &Product{
		Name:  "Silla",
		Price: 99.99,
		Stock: 10,
	}

	err := product.Validate()
	assert.NoError(t, err)
}
