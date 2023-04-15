package handler

import (
	"tugas-sesi12/database"
	"tugas-sesi12/repository/product_repository/product_pg"
	"tugas-sesi12/repository/user_repository/user_pg"
	"tugas-sesi12/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitializeDatabase()

	db := database.GetDatabaseInstance()

	productRepo := product_pg.NewProductPg(db)
	productService := service.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	userRepo := user_pg.NewUserPg(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, productRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	productRoute := route.Group("/products")
	{
		productRoute.GET("/", authService.Authentication(), productHandler.GetAllProducts)
		productRoute.GET("/:productId", authService.Authentication(), authService.Authorization(), productHandler.GetProductById)
		productRoute.POST("/", authService.Authentication(), productHandler.CreateProduct)
		productRoute.PUT("/:productId", authService.Authentication(), authService.Authorization(), productHandler.UpdateProductById)
	}

	route.Run(":" + port)
}
