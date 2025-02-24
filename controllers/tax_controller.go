package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type TaxController struct {
	taxService services.TaxService
}

func NewTaxController(service services.TaxService) *TaxController {
	return &TaxController{service}
}

func (c *TaxController) CreateTax(ctx *fiber.Ctx) error {
	var tax models.Tax
	if err := ctx.BodyParser(&tax); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.taxService.CreateTax(&tax); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(tax)
}

func (c *TaxController) GetAllTaxes(ctx *fiber.Ctx) error {
	taxes, err := c.taxService.GetAllTaxes()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(taxes)
}

func (c *TaxController) GetTaxByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	tax, err := c.taxService.GetTaxByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Tax not found"})
	}
	return ctx.JSON(tax)
}

func (c *TaxController) UpdateTax(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var tax models.Tax
	if err := ctx.BodyParser(&tax); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.taxService.UpdateTax(id, &tax); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(tax)
}

func (c *TaxController) DeleteTax(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.taxService.DeleteTax(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
