package http

import "github.com/gin-gonic/gin"

func NewRouter(handler *ProductHandler) *gin.Engine {
	router := gin.Default()

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
