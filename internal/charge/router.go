package charge

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gustagcosta/charge-friends/server/middleware"
)

func AddChargeRoutes(app *fiber.App, controller *ChargeController) {
	charges := app.Group("/charges", middleware.Protected)

	charges.Get("/", controller.GetAllCharges)
	charges.Post("/", controller.CreateNewCharge)
	charges.Delete("/:id", controller.DeleteCharge)
	charges.Put("/:id", controller.UpdateCharge)
}
