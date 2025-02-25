package routes

import (
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App, productController *controllers.ProductController,
	customerController *controllers.CustomerController,
	employeeController *controllers.EmployeeController,
	discountController *controllers.DiscountController,
	inventoryController *controllers.InventoryController,
	paymentController *controllers.PaymentController,
	receiptController *controllers.ReceiptController,
	taxController *controllers.TaxController,
	orderController *controllers.OrderController) {

	productGroup := app.Group("/products")
	productGroup.Post("/", productController.CreateProduct)
	productGroup.Post("/bulk", productController.CreateProductBulk)
	productGroup.Get("/", productController.GetAllProducts)
	productGroup.Get("/:id", productController.GetProductByID)
	productGroup.Put("/:id", productController.UpdateProduct)
	productGroup.Delete("/:id", productController.DeleteProduct)

	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customerController.CreateCustomer)
	customerGroup.Get("/", customerController.GetAllCustomers)
	customerGroup.Get("/:id", customerController.GetCustomerByID)
	customerGroup.Put("/:id", customerController.UpdateCustomer)
	customerGroup.Delete("/:id", customerController.DeleteCustomer)

	employeeGroup := app.Group("/employees")
	employeeGroup.Post("/", employeeController.CreateEmployee)
	employeeGroup.Get("/", employeeController.GetAllEmployees)
	employeeGroup.Get("/:id", employeeController.GetEmployeeByID)
	employeeGroup.Put("/:id", employeeController.UpdateEmployee)
	employeeGroup.Delete("/:id", employeeController.DeleteEmployee)

	discountGroup := app.Group("/discounts")
	discountGroup.Post("/", discountController.CreateDiscount)
	discountGroup.Get("/", discountController.GetAllDiscounts)
	discountGroup.Get("/:id", discountController.GetDiscountByID)
	discountGroup.Put("/:id", discountController.UpdateDiscount)
	discountGroup.Delete("/:id", discountController.DeleteDiscount)

	inventoryGroup := app.Group("/inventories")
	inventoryGroup.Post("/", inventoryController.CreateInventory)
	inventoryGroup.Get("/", inventoryController.GetAllInventories)
	inventoryGroup.Get("/:id", inventoryController.GetInventoryByID)
	inventoryGroup.Put("/:id", inventoryController.UpdateInventory)
	inventoryGroup.Delete("/:id", inventoryController.DeleteInventory)

	paymentGroup := app.Group("/payments")
	paymentGroup.Post("/", paymentController.CreatePayment)
	paymentGroup.Get("/", paymentController.GetAllPayments)
	paymentGroup.Get("/:id", paymentController.GetPaymentByID)
	paymentGroup.Put("/:id", paymentController.UpdatePayment)
	paymentGroup.Delete("/:id", paymentController.DeletePayment)

	receiptGroup := app.Group("/receipts")
	receiptGroup.Post("/", receiptController.CreateReceipt)
	receiptGroup.Get("/", receiptController.GetAllReceipts)
	receiptGroup.Get("/:id", receiptController.GetReceiptByID)
	receiptGroup.Put("/:id", receiptController.UpdateReceipt)
	receiptGroup.Delete("/:id", receiptController.DeleteReceipt)

	taxGroup := app.Group("/receipts")
	taxGroup.Post("/", taxController.CreateTax)
	taxGroup.Get("/", taxController.GetAllTaxes)
	taxGroup.Get("/:id", taxController.GetTaxByID)
	taxGroup.Put("/:id", taxController.UpdateTax)
	taxGroup.Delete("/:id", taxController.DeleteTax)

	orderGroup := app.Group("/orders")
	orderGroup.Post("/", orderController.CreateOrder)
	orderGroup.Get("/", orderController.GetAllOrders)
	orderGroup.Get("/:id", orderController.GetOrderByID)
	orderGroup.Put("/:id", orderController.UpdateOrder)
	orderGroup.Delete("/:id", orderController.DeleteOrder)
}
