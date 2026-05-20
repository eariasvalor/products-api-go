package http

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"proyectoGo/internal/adapters/http/middleware"
	"proyectoGo/internal/domain"
	"proyectoGo/internal/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter(handler *ProductHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
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

func TestHandler_GetAll_Success(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	expected := []*domain.Product{
		{ID: 1, Name: "Silla", Price: 99.99, Stock: 10},
	}
	service.On("GetAll").Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	service.AssertExpectations(t)
}

func TestHandler_GetAll_Error(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	service.On("GetAll").Return(nil, domain.NewInternalError("error interno"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	service.AssertExpectations(t)
}

func TestHandler_GetByID_Success(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	expected := &domain.Product{ID: 1, Name: "Silla", Price: 99.99, Stock: 10}
	service.On("GetByID", 1).Return(expected, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	service.AssertExpectations(t)
}

func TestHandler_GetByID_NotFound(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	service.On("GetByID", 99).Return(nil, domain.NewNotFoundError("producto no encontrado"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/products/99", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	service.AssertExpectations(t)
}

func TestHandler_Create_Success(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	input := &domain.Product{Name: "Silla", Description: "Silla de oficina", Price: 99.99, Stock: 10}
	expected := &domain.Product{ID: 1, Name: "Silla", Description: "Silla de oficina", Price: 99.99, Stock: 10}
	service.On("Create", input).Return(expected, nil)

	body, _ := json.Marshal(input)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	service.AssertExpectations(t)
}

func TestHandler_Delete_Success(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	service.On("Delete", 1).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/products/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	service.AssertExpectations(t)
}

func TestHandler_Delete_NotFound(t *testing.T) {
	service := new(mocks.ProductServiceMock)
	handler := NewProductHandler(service)
	router := setupRouter(handler)

	service.On("Delete", 99).Return(domain.NewNotFoundError("producto no encontrado"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/products/99", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	service.AssertExpectations(t)
}
