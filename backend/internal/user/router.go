package user

import (
	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(app *fiber.App, controller *UserController) {
	users := app.Group("/users")

	// add middlewares here

	// add routes here
	users.Post("/", controller.CreateNewUser)
	users.Post("/login", controller.Login)
}
