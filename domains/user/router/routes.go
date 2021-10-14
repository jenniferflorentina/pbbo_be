package router

import (
	"tubespbbo/domains/user/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteUsers(e *fiber.App) {
	e.Get("/users", controller.Find)
	e.Post("/users", controller.Create)

	//users := e.Group("/users")
	//{
	//	users.GET("", controller.Find)
	//	users.POST("", controller.Create)
	//}
}
