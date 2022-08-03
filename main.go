package main

import (
	"github.com/gofiber/fiber/v2"
	"golab.info/go-admin/database"
	"golab.info/go-admin/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")

}
