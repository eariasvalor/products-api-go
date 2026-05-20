package main

import (
	"log"
	"proyectoGo/config"
	httpAdapter "proyectoGo/internal/adapters/http"
	"proyectoGo/internal/adapters/postgres"
	"proyectoGo/internal/application"
)

func main() {
	cfg := config.Load()

	db, err := cfg.ConnectDB()
	if err != nil {
		log.Fatal("Error conectando a la base de datos: ", err)
	}
	defer db.Close()

	productRepository := postgres.NewProductRepository(db)
	productService := application.NewProductService(productRepository)
	productHandler := httpAdapter.NewProductHandler(productService)

	router := httpAdapter.NewRouter(productHandler)

	log.Println("Servidor corriendo en puerto:", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Error arrancando el servidor: ", err)
	}
}
