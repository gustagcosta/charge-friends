package client

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ClientController struct {
	storage *ClientStorage
}

func NewClientController(storage *ClientStorage) *ClientController {
	return &ClientController{
		storage: storage,
	}
}

func (c *ClientController) GetAllClients(ctx *fiber.Ctx) error {
	clients, err := c.storage.FindClients(int(ctx.Locals("userId").(float64)))
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed while fetching users",
		})
	}

	return ctx.JSON(clients)
}

func (c *ClientController) CreateNewClient(ctx *fiber.Ctx) error {
	req := new(ClientCreateRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "error parsing body",
		})
	}

	req.UserId = int(ctx.Locals("userId").(float64))

	err := req.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	clientId, err := c.storage.CreateNewClient(req, int(ctx.Locals("userId").(float64)))
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": clientId})
}
