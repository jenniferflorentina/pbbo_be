package router

import (
	"tubespbbo/domains/user/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteUsers(e *fiber.App) {
	e.Get("/users", controller.FindUser)
	e.Get("/users/:id", controller.FindOneUser)
	e.Post("/users", controller.CreateUser)
	e.Put("/users/:id", controller.UpdateUser)
	e.Delete("/users/:id", controller.DeleteUser)

	e.Post("/login", controller.Login)

	//users := e.Group("/users")
	//{
	//	users.GET("", controller.Find)
	//	users.POST("", controller.Create)
	//}
}
