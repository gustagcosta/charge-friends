package main

import (
	"api/internal/charge"
	"api/internal/client"
	"api/internal/user"
	"api/server/config"
	"api/server/queue"
	"api/server/storage"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// setup exit code for graceful shutdown
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// run the server
	cleanup, err := run(env)

	// run the cleanup after the server is terminated
	defer cleanup()

	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	// ensure the server is shutdown gracefully & app runs
	Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// start the server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// return a function to close the server and database
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	// init the storage
	db, err := storage.BootstrapPG(env.PG_URI)
	if err != nil {
		return nil, nil, err
	}

	// init the message broker
	sqsClient, err := queue.BootstrapSQS(env.QUEUE, env.AWS_ENDPOINT, env.AWS_REGION)
	if err != nil {
		return nil, nil, err
	}

	// create the fiber app
	app := fiber.New()

	// add middleware
	app.Use(cors.New())
	app.Use(logger.New())

	// add health check
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Healthy!")
	})

	userStorage := user.NewPostgresUserStorage(db)
	userController := user.NewUserController(userStorage)
	user.AddUserRoutes(app, userController)

	clientStorage := client.NewPostgresClientStorage(db)
	clientController := client.NewClientController(clientStorage)
	client.AddClientRoutes(app, clientController)

	chargeStorage := charge.NewPostgresChargeStorage(db)
	chargeController := charge.NewChargeController(chargeStorage, sqsClient)
	charge.AddChargeRoutes(app, chargeController)

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
