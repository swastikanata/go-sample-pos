package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type DiscountController struct {
	discountService services.DiscountService
}

func NewDiscountController(service services.DiscountService) *DiscountController {
	return &DiscountController{service}
}

func (c *DiscountController) CreateDiscount(ctx *fiber.Ctx) error {
	var discount models.Discount
	if err := ctx.BodyParser(&discount); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.discountService.CreateDiscount(&discount); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(discount)
}

func (c *DiscountController) GetAllDiscounts(ctx *fiber.Ctx) error {
	discounts, err := c.discountService.GetAllDiscounts()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(discounts)
}

func (c *DiscountController) GetDiscountByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	discount, err := c.discountService.GetDiscountByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Discount not found"})
	}
	return ctx.JSON(discount)
}

func (c *DiscountController) UpdateDiscount(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var discount models.Discount
	if err := ctx.BodyParser(&discount); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.discountService.UpdateDiscount(id, &discount); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(discount)
}

func (c *DiscountController) DeleteDiscount(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.discountService.DeleteDiscount(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
