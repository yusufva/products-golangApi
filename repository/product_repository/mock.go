package product_repository

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
)

var (
	CreateProduct        func(productPayload *entity.Product) (*entity.Product, errrs.MessageErr)
	GetProductById       func(productId int) (*entity.Product, errrs.MessageErr)
	UpdateProductById    func(payload entity.Product) errrs.MessageErr
	GetAllProducts       func() ([]*entity.Product, errrs.MessageErr)
	GetAllProductsByUser func(userId int) ([]*entity.Product, errrs.MessageErr)
)

type productRepoMock struct{}

func NewProductRepoMock() ProductRepository {
	return &productRepoMock{}
}

func (p *productRepoMock) CreateProduct(productPayload *entity.Product) (*entity.Product, errrs.MessageErr) {
	return CreateProduct(productPayload)
}
func (p *productRepoMock) GetProductById(productId int) (*entity.Product, errrs.MessageErr) {
	return GetProductById(productId)
}
func (p *productRepoMock) UpdateProductById(payload entity.Product) errrs.MessageErr {
	return UpdateProductById(payload)
}
func (p *productRepoMock) GetAllProducts() ([]*entity.Product, errrs.MessageErr) {
	return GetAllProducts()
}

func (p *productRepoMock) GetAllProductsByUser(userId int) ([]*entity.Product, errrs.MessageErr) {
	return GetAllProductsByUser(userId)
}
