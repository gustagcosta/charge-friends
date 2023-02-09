package main

import (
	"backend/internal/storage"
	"backend/internal/user"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	cleanup, err := run()
	defer cleanup()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	Gracefully()
}

func run() (func(), error) {
	app, cleanup, err := buildServer()
	if err != nil {
		return nil, err
	}

	go func() {
		app.Listen("0.0.0.0:" + os.Getenv("PORT"))
	}()

	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer() (*fiber.App, func(), error) {
	db, err := storage.BootstrapPG(os.Getenv("PG_URI"))
	if err != nil {
		return nil, nil, err
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	userStore := user.NewUserStorage(db)
	userController := user.NewUserController(userStore)
	user.AddUserRoutes(app, userController)

	return app, func() {
		storage.ClosePG(db)
	}, nil
}

func Gracefully() {
	quit := make(chan os.Signal, 1)
	defer close(quit)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
