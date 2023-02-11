package client

import (
	"backend/internal/shared"

	"github.com/gofiber/fiber/v2"
)

func AddClientRoutes(app *fiber.App, controller *ClientController) {
	clients := app.Group("/clients", shared.Protected)

	clients.Get("/", controller.GetAllClients)
	clients.Post("/", controller.CreateNewClient)
	clients.Delete("/:id", controller.DeleteClient)
	clients.Put("/:id", controller.UpdateClient)
}
