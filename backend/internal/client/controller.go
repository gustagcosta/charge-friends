package client

import (
	"fmt"
	"strconv"

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
			"message": "failed while fetching clients",
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
			"message": "failed at register client",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": clientId})
}

func (c *ClientController) DeleteClient(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	err = c.storage.DeleteClientByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete client",
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *ClientController) UpdateClient(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	req := new(ClientCreateRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "error parsing body",
		})
	}

	client, err := c.storage.FindClientByID(id)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "client not found",
		})
	}

	if req.Email == "" {
		req.Email = client.Email
	}

	if req.Whatsapp == "" {
		req.Whatsapp = client.Whatsapp
	}

	if req.Name == "" {
		req.Name = client.Name
	}

	err = req.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = c.storage.UpdateClient(id, req)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at update client",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
