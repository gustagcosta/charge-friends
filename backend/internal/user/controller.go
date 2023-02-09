package user

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	storage *UserStorage
}

func NewUserController(storage *UserStorage) *UserController {
	return &UserController{
		storage: storage,
	}
}

func (c *UserController) CreateNewUser(ctx *fiber.Ctx) error {
	req := new(UserCreateRequest)
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

	user, err := c.storage.FindUserByEmail(req.Email)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	if len(user.Email) != 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "user already exists",
		})
	}

	bytes, err := HashPassword(user.Password)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	req.Password = bytes

	err = c.storage.CreateNewUser(*req, ctx.Context())
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	return ctx.SendStatus(http.StatusCreated)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	req := &UserLoginRequest{}
	if err := ctx.BodyParser(req); err != nil {
		fmt.Println("error: ", err)
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

	user, err := c.storage.FindUserByEmail(req.Email)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	if len(user.Email) == 0 {
		return ctx.Status(fiber.StatusMethodNotAllowed).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	fmt.Println("hash", user.Password)
	fmt.Println("password", req.Password)

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))

	if passwordMatch != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "password not match",
		})
	}

	return ctx.SendStatus(http.StatusOK)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
