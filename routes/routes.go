package routes

import (
	"github.com/gofiber/fiber/v2"
	"golab.info/go-admin/controllers"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
	app.Get("/other", controllers.Other)
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

}
