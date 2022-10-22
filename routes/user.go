package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harikesh-yadav/gofiber_api/controllers"
)

func User(app *fiber.App) {

	route := app.Group("/user")
	route.Post("/registration", controllers.Registration)
	route.Post("/login", controllers.Login)
}
