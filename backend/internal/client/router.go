package client

import (
	"backend/internal/shared"

	"github.com/gofiber/fiber/v2"
)

func AddClientRoutes(app *fiber.App, controller *ClientController) {
	clients := app.Group("/clients")

	clients.Get("/", shared.Protected, controller.GetAllClients)
	clients.Post("/", shared.Protected, controller.CreateNewClient)
}
