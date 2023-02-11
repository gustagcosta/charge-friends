package user

import (
	"backend/internal/shared"

	"github.com/gofiber/fiber/v2"
)

func AddUserRoutes(app *fiber.App, controller *UserController) {
	users := app.Group("/users")

	users.Post("/", controller.CreateNewUser)
	users.Post("/login", controller.Login)
	users.Get("/private", shared.Protected, controller.Private)
}
