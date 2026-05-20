package domain

import "errors"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required,min=2,max=100"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"min=0"`
}

func NewProduct(name string, description string, price float64, stock int) (*Product, error) {
	p := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("el nombre es obligatorio")
	}
	if p.Price <= 0 {
		return errors.New("el precio debe ser mayor a 0")
	}
	if p.Stock < 0 {
		return errors.New("el stock no puede ser negativo")
	}
	return nil
}
