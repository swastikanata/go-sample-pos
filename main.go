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

	customerRepo := repositories.NewCustomerRepository(database.DB)
	customerService := services.NewCustomerService(customerRepo)
	customerController := controllers.NewCustomerController(customerService)

	employeeRepo := repositories.NewEmployeeRepository(database.DB)
	employeeService := services.NewEmployeeService(employeeRepo)
	employeeController := controllers.NewEmployeeController(employeeService)

	discountRepo := repositories.NewDiscountRepository(database.DB)
	discountService := services.NewDiscountService(discountRepo)
	discountController := controllers.NewDiscountController(discountService)

	inventoryRepo := repositories.NewInventoryRepository(database.DB)
	inventoryService := services.NewInventoryService(inventoryRepo)
	inventoryController := controllers.NewInventoryController(inventoryService)

	paymentRepo := repositories.NewPaymentRepository(database.DB)
	paymentService := services.NewPaymentService(paymentRepo)
	paymentController := controllers.NewPaymentController(paymentService)

	receiptRepo := repositories.NewReceiptRepository(database.DB)
	receiptService := services.NewReceiptService(receiptRepo)
	receiptController := controllers.NewReceiptController(receiptService)

	taxRepo := repositories.NewTaxRepository(database.DB)
	taxService := services.NewTaxService(taxRepo)
	taxController := controllers.NewTaxController(taxService)

	orderRepo := repositories.NewOrderRepository(database.DB)
	orderService := services.NewOrderService(orderRepo)
	orderController := controllers.NewOrderController(orderService)

	// Setup Routes
	routes.SetupRoutes(app, productController, customerController,
		employeeController, discountController, inventoryController,
		paymentController, receiptController, taxController, orderController)

	// Start server
	err := app.Listen(":8080")
	if err != nil {
		fmt.Println(err)
	}
}
