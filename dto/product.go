package dto

import (
	"time"
)

type NewProductRequest struct {
	Title       string `json:"title" valid:"required~title cannot be empty"`
	Description string `json:"description" valid:"required~description cannot be empty"`
	UserId      int    `json:"userId" valid:"int,required~userId cannot be empty"`
}

type NewProductResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type ProductResponse struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	UserId      int       `json:"userId"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}

type GetProductsResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Data       []ProductResponse `json:"data"`
}
