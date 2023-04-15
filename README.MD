# Chapter 3 Session 2
## Challenge
Making simple CRUD API with middleware, multilevel users, and authorization product by user id using Golang

### Summary
Simple CRUD with multilevel Users, middleware, and authorization for Create and Update Product. There is two different user level in this API, first user level is ```Admin``` and the second user level is ```Customer```

```Admin``` user level can modify all of the products that are in database, meanwhile ```customer``` user level can only modify his product.

### Technology used
* [Golang](https://go.dev/)
* DDD structure
* [PostgreSQL](https://www.postgresql.org/) 
* [GORM](https://gorm.io/)
* [Gin Framework](https://gin-gonic.com/)
* [JWT](https://jwt.io/)

### endpoints
```POST /users/register```      => Register a new user 
```POST /users/login```         => Login as registered user
```GET /products/```            => Get all products based on user level
```GET /products/:productId```  => Get product based on id with user level authorization
```POST /products/```           => Creating New Product data
```PUT /products/:productId```  => Update product based on id with user level authorization