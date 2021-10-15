package router

import (
	"tubespbbo/domains/payment/controller"

	"github.com/gofiber/fiber/v2"
)

func RoutePayments(e *fiber.App) {
	e.Get("/payments", controller.FindPayment)
	e.Get("/payments/:id", controller.FindOnePayment)
	e.Post("/payments", controller.CreatePayment)
}

func RoutePaymentMethods(e *fiber.App) {
	e.Get("/payment-methods", controller.FindPaymentMethod)
	e.Get("/payment-methods/:id", controller.FindOnePaymentMethod)
	e.Post("/payment-methods", controller.CreatePaymentMethod)
}
