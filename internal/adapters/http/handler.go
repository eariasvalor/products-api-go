package http

import (
	"net/http"
	"proyectoGo/internal/adapters/http/middleware"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/ports/input"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service input.ProductService
}

func NewProductHandler(service input.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	products, err := h.service.GetAll()
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domain.NewBadRequestError("id inválido"))
		return
	}

	product, err := h.service.GetByID(id)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.Error(domain.NewBadRequestError(middleware.TranslateValidationErrors(err)))
		return
	}

	created, err := h.service.Create(&product)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domain.NewBadRequestError("id inválido"))
		return
	}

	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.Error(domain.NewBadRequestError(middleware.TranslateValidationErrors(err)))
		return
	}

	updated, err := h.service.Update(id, &product)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(domain.NewBadRequestError("id inválido"))
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "producto eliminado"})
}
