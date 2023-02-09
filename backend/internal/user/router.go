package user

import (
	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(app *fiber.App, controller *UserController) {
	users := app.Group("/users")

	users.Post("/", controller.CreateNewUser)
	users.Post("/login", controller.Login)
	users.Get("/private", Protected, controller.Private)
}
