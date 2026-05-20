package domain

import "errors"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
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
