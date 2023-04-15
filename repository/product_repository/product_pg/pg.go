package product_pg

import (
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/repository/product_repository"

	"gorm.io/gorm"
)

type productPg struct {
	db *gorm.DB
}

func NewProductPg(db *gorm.DB) product_repository.ProductRepository {
	return &productPg{
		db: db,
	}
}

func (m *productPg) CreateProduct(productPayload *entity.Product) (*entity.Product, errrs.MessageErr) {
	result := m.db.Create(productPayload)

	if result.Error != nil {
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	row := result.Row()

	var product entity.Product
	row.Scan(row, &product)

	return &product, nil
}

func (m *productPg) GetProductById(productId int) (*entity.Product, errrs.MessageErr) {
	var product entity.Product
	result := m.db.First(&product, productId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errrs.NewNotFoundError("product not found")
		}
		return nil, errrs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}
func (m *productPg) UpdateProductById(payload entity.Product) errrs.MessageErr {
	err := m.db.Model(payload).Updates(entity.Product{Title: payload.Title, Description: payload.Description}).Error

	if err != nil {
		return errrs.NewInternalServerError("error while saving product")
	}

	return nil
}

func (m *productPg) GetAllProducts() ([]*entity.Product, errrs.MessageErr) {
	var products []*entity.Product

	err := m.db.Find(&products).Error

	if err != nil {
		return nil, errrs.NewInternalServerError("error getting data")
	}

	return products, nil
}

func (m *productPg) GetAllProductsByUser(userId int) ([]*entity.Product, errrs.MessageErr) {
	var products []*entity.Product

	err := m.db.Find(&products, "user_id = ?", userId).Error

	if err != nil {
		return nil, errrs.NewInternalServerError("error getting data")
	}

	return products, nil
}
