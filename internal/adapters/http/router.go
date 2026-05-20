package http

import (
	"proyectoGo/internal/adapters/http/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *ProductHandler) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.ErrorHandler())

	products := router.Group("/products")
	{
		products.GET("", handler.GetAll)
		products.GET("/:id", handler.GetByID)
		products.POST("", handler.Create)
		products.PUT("/:id", handler.Update)
		products.DELETE("/:id", handler.Delete)
	}

	return router
}
