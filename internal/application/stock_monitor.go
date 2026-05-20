package application

import (
	"log"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/ports/output"
	"time"
)

type StockMonitor struct {
	repo      output.ProductRepository
	threshold int
}

func NewStockMonitor(repo output.ProductRepository, threshold int) *StockMonitor {
	return &StockMonitor{repo: repo, threshold: threshold}
}

func (m *StockMonitor) CheckLowStock() []*domain.Product {
	products, err := m.repo.FindAll()
	if err != nil {
		return nil
	}

	var lowStock []*domain.Product
	for _, p := range products {
		if p.Stock < m.threshold {
			lowStock = append(lowStock, p)
		}
	}
	return lowStock
}

func (m *StockMonitor) Start(interval time.Duration) func() {
	stop := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				lowStock := m.CheckLowStock()
				for _, p := range lowStock {
					log.Printf("⚠️  Stock bajo: %s tiene %d unidades", p.Name, p.Stock)
				}
			case <-stop:
				log.Println("Stock monitor parado")
				return
			}
		}
	}()

	return func() { close(stop) }
}
