package user

import (
	"fmt"
	"time"

	"api/server/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	storage UserStorage
}

func NewUserController(storage UserStorage) *UserController {
	return &UserController{
		storage: storage,
	}
}

type Claims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
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

	user, err := c.storage.FindByEmail(req.Email)
	if err != nil {
		fmt.Println("error: ", err) // improve this
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	if len(user.Email) != 0 {
		return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "user already exists",
		})
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	encryptedPassword := string(bytes)
	req.Password = encryptedPassword

	err = c.storage.Create(req)
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	req := &UserLoginRequest{}
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

	user, err := c.storage.FindByEmail(req.Email)
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

	passwordMatch := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if passwordMatch != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "password not match",
		})
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	env, err := config.LoadConfig()
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(env.JWT_KEY))
	if err != nil {
		fmt.Println("error: ", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "failed at register user",
		})
	}

	return ctx.JSON(fiber.Map{"token": tokenString})
}

func (c *UserController) Private(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{"id": ctx.Locals("userId")})
}
