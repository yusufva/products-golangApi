package handler

import (
	"net/http"
	"tugas-sesi12/dto"
	"tugas-sesi12/entity"
	"tugas-sesi12/pkg/errrs"
	"tugas-sesi12/pkg/helpers"
	"tugas-sesi12/service"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{
		productService: productService,
	}
}

func (p productHandler) GetProductById(c *gin.Context) {
	productId, err := helpers.GetParamsId(c, "productId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := p.productService.GetProductById(productId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, &response)
}

func (p productHandler) GetAllProducts(c *gin.Context) {
	// products := []entity.Product{}

	user := c.MustGet("userData").(entity.User)

	if user.Level != entity.Admin {
		allProductsUser, err := p.productService.GetAllProductsByUser(user.Id)
		if err != nil {
			c.JSON(err.Status(), err)
		}
		c.JSON(http.StatusOK, allProductsUser)
		return
	}

	allProducts, err := p.productService.GetAllProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, allProducts)
}

func (p productHandler) CreateProduct(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	user := c.MustGet("userData").(entity.User)

	newProduct, err := p.productService.CreateProduct(user.Id, productRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, newProduct)
}

func (p productHandler) UpdateProductById(c *gin.Context) {
	var productRequest dto.NewProductRequest

	if err := c.ShouldBindJSON(&productRequest); err != nil {
		errBindJson := errrs.NewUnprocessibleEntityError("invalid request body")

		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	productId, err := helpers.GetParamsId(c, "productId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	// user := c.MustGet("userData").(entity.User)

	response, err := p.productService.UpdateProductById(productId, productRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
