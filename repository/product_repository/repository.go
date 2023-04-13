package product_repository

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
)

type ProductRepository interface {
	CreateProduct(productPayload *entity.Product) (*entity.Product, errrs.MessageErr)
	GetProductById(productId int) (*entity.Product, errrs.MessageErr)
	UpdateProductById(payload entity.Product) errrs.MessageErr
	GetAllProducts() ([]*entity.Product, errrs.MessageErr)
}
