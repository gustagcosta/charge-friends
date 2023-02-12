package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustagcosta/charge-friends/server/middleware"
)

func AddUserRoutes(app *fiber.App, controller *UserController) {
	users := app.Group("/users")

	users.Post("/", controller.CreateNewUser)
	users.Post("/login", controller.Login)
	users.Get("/private", middleware.Protected, controller.Private)
}
