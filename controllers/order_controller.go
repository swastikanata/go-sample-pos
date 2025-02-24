package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService services.OrderService
}

func NewOrderController(service services.OrderService) *OrderController {
	return &OrderController{service}
}

func (c *OrderController) CreateOrder(ctx *fiber.Ctx) error {
	var order models.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.orderService.CreateOrder(&order); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(order)
}

func (c *OrderController) GetAllOrders(ctx *fiber.Ctx) error {
	orders, err := c.orderService.GetAllOrders()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(orders)
}

func (c *OrderController) GetOrderByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	order, err := c.orderService.GetOrderByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Order not found"})
	}
	return ctx.JSON(order)
}

func (c *OrderController) UpdateOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var order models.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.orderService.UpdateOrder(id, &order); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(order)
}

func (c *OrderController) DeleteOrder(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.orderService.DeleteOrder(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
