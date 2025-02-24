package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type InventoryController struct {
	inventoryService services.InventoryService
}

func NewInventoryController(service services.InventoryService) *InventoryController {
	return &InventoryController{service}
}

func (c *InventoryController) CreateInventory(ctx *fiber.Ctx) error {
	var Inventory models.Inventory
	if err := ctx.BodyParser(&Inventory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.inventoryService.CreateInventory(&Inventory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(Inventory)
}

func (c *InventoryController) GetAllInventories(ctx *fiber.Ctx) error {
	inventories, err := c.inventoryService.GetAllInventories()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(inventories)
}

func (c *InventoryController) GetInventoryByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	Inventory, err := c.inventoryService.GetInventoryByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Inventory not found"})
	}
	return ctx.JSON(Inventory)
}

func (c *InventoryController) UpdateInventory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var Inventory models.Inventory
	if err := ctx.BodyParser(&Inventory); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.inventoryService.UpdateInventory(id, &Inventory); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(Inventory)
}

func (c *InventoryController) DeleteInventory(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.inventoryService.DeleteInventory(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
