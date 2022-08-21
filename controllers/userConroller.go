package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golab.info/go-admin/database"
	"golab.info/go-admin/models"
)

func AllUsers(c *fiber.Ctx) error {
	limit := 5
	var users []models.User

	database.DB.Preload("Role").Limit(limit).Find(&users)

	//return c.JSON(users)
	return c.JSON(fiber.Map{
		"data": users,
		"meta": fiber.Map{
			"page": 1,
		},
	})
}

func CraeteUser(c *fiber.Ctx) error {

	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)

}

func GetUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Preload("Role").Find(&user)

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)

}

func DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	user := models.User{
		Id: uint(id),
	}

	database.DB.Delete(&user)

	return nil

}
