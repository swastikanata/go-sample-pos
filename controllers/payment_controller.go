package controllers

import (
	"github.com/aronipurwanto/go-sample-pos/models"
	"github.com/aronipurwanto/go-sample-pos/services"
	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	paymentService services.PaymentService
}

func NewPaymentController(service services.PaymentService) *PaymentController {
	return &PaymentController{service}
}

func (c *PaymentController) CreatePayment(ctx *fiber.Ctx) error {
	var payment models.Payment
	if err := ctx.BodyParser(&payment); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.paymentService.CreatePayment(&payment); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(201).JSON(payment)
}

func (c *PaymentController) GetAllPayments(ctx *fiber.Ctx) error {
	payments, err := c.paymentService.GetAllPayments()
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(payments)
}

func (c *PaymentController) GetPaymentByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	payment, err := c.paymentService.GetPaymentByID(id)
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{"error": "Payment not found"})
	}
	return ctx.JSON(payment)
}

func (c *PaymentController) UpdatePayment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var payment models.Payment
	if err := ctx.BodyParser(&payment); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "Invalid request"})
	}
	if err := c.paymentService.UpdatePayment(id, &payment); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(payment)
}

func (c *PaymentController) DeletePayment(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if err := c.paymentService.DeletePayment(id); err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(204).Send(nil)
}
