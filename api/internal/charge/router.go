package charge

import (
	"api/server/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddChargeRoutes(app *fiber.App, controller *ChargeController) {
	charges := app.Group("/charges", middleware.Protected)

	charges.Get("/", controller.GetAllCharges)
	charges.Post("/", controller.CreateNewCharge)
	charges.Delete("/:id", controller.DeleteCharge)
	charges.Put("/:id", controller.UpdateCharge)
	charges.Post("notification/:id", controller.Notification)
}
