package service

import (
	"net/http"
	"tugas-sesi12/dto"
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/pkg/helpers"
	"tugas-sesi12/repository/product_repository"
)

type ProductService interface {
	CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errrs.MessageErr)
	UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errrs.MessageErr)
	GetProductById(productId int) (*dto.ProductResponse, errrs.MessageErr)
	GetAllProducts() (*dto.GetProductsResponse, errrs.MessageErr)
	GetAllProductsByUser(userId int) (*dto.GetProductsResponse, errrs.MessageErr)
}

type productService struct {
	productRepo product_repository.ProductRepository
}

func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errrs.MessageErr) {
	productRequest := &entity.Product{
		Title:       payload.Title,
		Description: payload.Description,
		UserId:      userId,
	}

	_, err := p.productRepo.CreateProduct(productRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "product has been successfully created",
	}

	return &response, nil
}

func (p *productService) UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errrs.MessageErr) {
	err := helpers.ValidateStruct(productRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Product{
		Id:          productId,
		Title:       productRequest.Title,
		Description: productRequest.Description,
		UserId:      productRequest.UserId,
	}

	err = p.productRepo.UpdateProductById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "product data successfully updated",
	}

	return &response, nil
}

func (p *productService) GetProductById(productId int) (*dto.ProductResponse, errrs.MessageErr) {
	result, err := p.productRepo.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	response := result.EntityToProductResponseDto()

	return &response, nil
}

func (p *productService) GetAllProducts() (*dto.GetProductsResponse, errrs.MessageErr) {
	products, err := p.productRepo.GetAllProducts()

	if err != nil {
		return nil, err
	}

	productResponse := []dto.ProductResponse{}

	for _, eachProduct := range products {
		productResponse = append(productResponse, eachProduct.EntityToProductResponseDto())
	}

	response := dto.GetProductsResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "product data have been sent successfully",
		Data:       productResponse,
	}

	return &response, nil
}

func (p *productService) GetAllProductsByUser(userId int) (*dto.GetProductsResponse, errrs.MessageErr) {
	products, err := p.productRepo.GetAllProductsByUser(userId)

	if err != nil {
		return nil, err
	}

	productResponse := []dto.ProductResponse{}

	for _, eachProduct := range products {
		productResponse = append(productResponse, eachProduct.EntityToProductResponseDto())
	}

	response := dto.GetProductsResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "product data have been sent successfully",
		Data:       productResponse,
	}

	return &response, nil
}
