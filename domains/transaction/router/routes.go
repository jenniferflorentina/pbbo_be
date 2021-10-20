package router

import (
	"tubespbbo/domains/transaction/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteTransactions(e *fiber.App) {
	e.Get("/transactions", controller.FindTransaction)
	e.Get("/transactions/:id", controller.FindOneTransaction)
	e.Post("/transactions", controller.CreateTransaction)
	e.Put("/transactions/:id", controller.UpdateTransaction)
	e.Delete("/transactions/:id", controller.DeleteTransaction)
}

func RouteTransactionMethods(e *fiber.App) {
	e.Get("/transaction-details", controller.FindTransactionDetail)
	e.Get("/transaction-details/:id", controller.FindOneTransactionDetail)
	e.Post("/transaction-details", controller.CreateTransactionDetail)
	e.Put("/transaction-details/:id", controller.UpdateTransactionDetail)
	e.Delete("/transaction-details/:id", controller.DeleteTransactionDetail)
}
