package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"golab.info/go-admin/database"
	"golab.info/go-admin/models"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

type RoleDTO struct {
	Name        string
	Permissions []float64
}

func CraeteRole(c *fiber.Ctx) error {
	var roleDTO RoleDTO

	err := c.BodyParser(&roleDTO)

	if err != nil {
		return err
	}

	list := roleDTO.Permissions

	permissions := make([]models.Permission, len(list))

	for i, permissionId := range list {
		permissions[i] = models.Permission{
			Id: uint(permissionId),
		}
	}
	
	role := models.Role{
		Name:        roleDTO.Name,
		Permissions: permissions,
	}

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Find(&role)

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)

}

func DeleteRole(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	role := models.Role{
		Id: uint(id),
	}

	database.DB.Delete(&role)

	return nil

}
