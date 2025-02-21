package main

import (
	"fmt"
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/aronipurwanto/go-sample-pos/database"
	"github.com/aronipurwanto/go-sample-pos/repositories"
	"github.com/aronipurwanto/go-sample-pos/routes"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to database
	database.ConnectDatabase()

	// Dependency Injection
	productRepo := repositories.NewProductRepository(database.DB)
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	// Setup Routes
	routes.SetupRoutes(app, productController)

	// Start server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
