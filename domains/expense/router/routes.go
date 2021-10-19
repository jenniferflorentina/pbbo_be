package router

import (
	"tubespbbo/domains/expense/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteExpenses(e *fiber.App) {
	e.Get("/expenses", controller.FindExpense)
	e.Get("/expenses/:id", controller.FindOneExpense)
	e.Post("/expenses", controller.CreateExpense)
	e.Put("/expenses/:id", controller.UpdateExpense)
	e.Delete("/expenses/:id", controller.DeleteExpense)
}