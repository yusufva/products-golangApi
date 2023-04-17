package service

import (
	"net/http"
	"testing"
	"time"
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/repository/product_repository"

	"github.com/stretchr/testify/assert"
)

func Test_ProductService_GetProductById_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)
	currentTime := time.Now()

	product_repository.GetProductById = func(productId int) (*entity.Product, errrs.MessageErr) {
		return &entity.Product{
			Id:          1,
			Title:       "Test Product",
			Description: "Test Product Description",
			UserId:      1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		}, nil
	}

	result, err := productService.GetProductById(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "Test Product", result.Title)
	assert.Equal(t, "Test Product Description", result.Description)
	assert.Equal(t, 1, result.UserId)
}

func Test_ProductService_GetProductById_NotFoundError(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)

	product_repository.GetProductById = func(productId int) (*entity.Product, errrs.MessageErr) {
		return nil, errrs.NewNotFoundError("product data not found")
	}

	response, err := productService.GetProductById(1)

	assert.Nil(t, response)
	assert.NotNil(t, err)
	assert.Equal(t, http.StatusNotFound, err.Status())
	assert.Equal(t, "product data not found", err.Message())
	assert.Equal(t, "NOT_FOUND", err.Error())
}

func Test_ProductService_GetAllProducts_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)
	currentTime := time.Now()

	products := []*entity.Product{
		{
			Id:          1,
			Title:       "Test Product 1",
			Description: "Product 1 Test Description",
			UserId:      1,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		},
	}

	product_repository.GetAllProducts = func() ([]*entity.Product, errrs.MessageErr) {
		return products, nil
	}

	respponse, err := productService.GetAllProducts()

	assert.Nil(t, err)
	assert.NotNil(t, respponse)
	assert.Equal(t, 1, len(respponse.Data))
	assert.Equal(t, "Test Product 1", respponse.Data[0].Title)
	assert.Equal(t, http.StatusOK, respponse.StatusCode)
	assert.Equal(t, "product data have been sent successfully", respponse.Message)
}

func Test_ProductService_GetAllProducts_NotFoundError(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)

	product_repository.GetAllProducts = func() ([]*entity.Product, errrs.MessageErr) {
		return []*entity.Product{}, nil
	}

	response, err := productService.GetAllProducts()

	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, 0, len(response.Data))
}
