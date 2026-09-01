package routes

import (
	"gocrud/pkg/controllers"
	"gocrud/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {

	app.Get("/", controllers.Health)

	//Get all Users
	app.Get("/users", middleware.SayHello, controllers.GetAllUser)

	//Create User
	app.Post("/users/create", controllers.CreateUser)

	//update User
	app.Put("/users/update/:id", controllers.UpdateUser)

	//delete User
	app.Delete("/users/delete/:id", controllers.DeleteUser)
}
