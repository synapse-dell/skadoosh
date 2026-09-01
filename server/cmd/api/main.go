package main

import (
	"gocrud/pkg/middleware"
	"gocrud/pkg/routes"
	"log"

	"gocrud/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Load()
	log.Println("DB_NAME", config.Get("DB_NAME"))
	config.ConnectDatabase()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Something went wrong",
			})
		},
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
	}))
	app.Use(middleware.Logger)
	routes.SetupUserRoutes(app)

	port := config.Get("PORT")
	app.Listen(":" + port)
}
