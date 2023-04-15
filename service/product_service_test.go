package service

import (
	"testing"
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/repository/product_repository"

	"github.com/stretchr/testify/assert"
)

func Test_ProductService_GetProductById_Success(t *testing.T) {
	productRepo := product_repository.NewProductRepoMock()
	productService := NewProductService(productRepo)

	product_repository.GetProductById = func(productId int) (*entity.Product, errrs.MessageErr) {
		return &entity.Product{
			Id:          1,
			Title:       "buku belajar",
			Description: "ini buku",
			UserId:      1,
		}, nil
	}

	result, err := productService.GetProductById(1)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.Id)
	assert.Equal(t, "buku belajar", result.Title)
	assert.Equal(t, "ini buku", result.Description)
	assert.Equal(t, 1, result.UserId)
}
