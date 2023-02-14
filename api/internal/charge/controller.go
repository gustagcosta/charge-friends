package charge

import (
	"api/server/queue"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ChargeController struct {
	storage   ChargeStorage
	sqsClient *queue.SqsClient
}

func NewChargeController(storage ChargeStorage, sqsClient *queue.SqsClient) *ChargeController {
	return &ChargeController{
		storage:   storage,
		sqsClient: sqsClient,
	}
}

func (c *ChargeController) GetAllCharges(ctx *fiber.Ctx) error {
	charges, err := c.storage.FindByUserID(int(ctx.Locals("userId").(float64)))
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed while fetching charges",
		})
	}

	return ctx.JSON(charges)
}

func (c *ChargeController) CreateNewCharge(ctx *fiber.Ctx) error {
	req := new(CreateChargeRequest)
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

	chargeId, err := c.storage.Create(req, int(ctx.Locals("userId").(float64)))
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register charge",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"id": chargeId})
}

func (c *ChargeController) DeleteCharge(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	err = c.storage.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed to delete charge",
		})
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (c *ChargeController) UpdateCharge(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	req := new(CreateChargeRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": "error parsing body",
		})
	}

	charge, err := c.storage.FindByID(id)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "charge not found",
		})
	}

	if req.Value == 0 {
		req.Value = charge.Value
	}

	if req.NotificationDate.IsZero() {
		req.NotificationDate = charge.NotificationDate
	}

	if req.Observation == "" {
		req.Observation = charge.Observation
	}

	if req.ClientId == 0 {
		req.ClientId = charge.ClientId
	}

	err = req.Validate()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = c.storage.Update(id, req)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at update charge",
		})
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *ChargeController) Notification(ctx *fiber.Ctx) error {
	messageId, err := c.sqsClient.SendMessage()
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at send message",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": messageId,
	})
}
