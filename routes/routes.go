package routes

import (
	"github.com/aronipurwanto/go-sample-pos/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, controller *controllers.ProductController) {
	productGroup := app.Group("/products")
	productGroup.Post("/", controller.CreateProduct)
	productGroup.Get("/", controller.GetAllProducts)
	productGroup.Get("/:id", controller.GetProductByID)
	productGroup.Put("/:id", controller.UpdateProduct)
	productGroup.Delete("/:id", controller.DeleteProduct)
}
