package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golab.info/go-admin/database"
	"golab.info/go-admin/models"
)

func AllProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))

	return c.JSON(models.Paginate(database.DB, &models.Product{}, page))
}
