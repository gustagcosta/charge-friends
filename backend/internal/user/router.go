package user

import "github.com/gofiber/fiber/v2"

func AddUserRoutes(app *fiber.App, controller *UserController) {
	users := app.Group("/users")

	// add middlewares here

	// add routes here
	users.Get("/", controller.getAll)
}
