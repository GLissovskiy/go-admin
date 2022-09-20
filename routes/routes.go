package routes

import (
	"github.com/gofiber/fiber/v2"
	"golab.info/go-admin/controllers"
	"golab.info/go-admin/middlewares"
)

func Setup(app *fiber.App) {

	app.Get("/", controllers.Hello)
	app.Get("/other", controllers.Other)
	app.Post("api/register", controllers.Register)
	app.Post("api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Put("api/users/info", controllers.UpadteInfo)
	app.Put("api/users/password", controllers.UpadtePassword)

	app.Get("api/user", controllers.User)
	app.Post("api/logout", controllers.Logout)

	app.Get("api/users", controllers.AllUsers)
	app.Post("api/users", controllers.CraeteUser)
	app.Get("api/users/:id", controllers.GetUser)
	app.Put("api/users/:id", controllers.UpdateUser)
	app.Delete("api/users/:id", controllers.DeleteUser)

	app.Get("api/roles", controllers.AllRoles)
	app.Post("api/roles", controllers.CraeteRole)
	app.Get("api/roles/:id", controllers.GetRole)
	app.Put("api/roles/:id", controllers.UpdateRole)
	app.Delete("api/roles/:id", controllers.DeleteRole)

	app.Get("api/permissions", controllers.AllPermissions)

	app.Get("api/products", controllers.AllProducts)
	app.Post("api/products", controllers.CraeteProduct)
	app.Get("api/products/:id", controllers.GetProduct)
	app.Put("api/products/:id", controllers.UpdateProduct)
	app.Delete("api/products/:id", controllers.DeleteProduct)

	app.Post("api/upload", controllers.Upload)
	app.Static("api/uploads", "./uploads")

	app.Get("api/orders", controllers.AllOrders)
	app.Post("api/export", controllers.Export)
	app.Get("api/chart", controllers.Chart)
}
