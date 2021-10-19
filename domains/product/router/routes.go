package router

import (
	"tubespbbo/domains/pruduct/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteProducts(e *fiber.App) {
	e.Get("/products", controller.FindProduct)
	e.Get("/products/:id", controller.FindOneProduct)
	e.Post("/products", controller.CreateProduct)
	e.Put("/products/:id", controller.UpdateProduct)
	e.Delete("/products/:id", controller.DeleteProduct)
}