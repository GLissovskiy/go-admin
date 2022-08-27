package models

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Paginate(db *gorm.DB, page int) fiber.Map {
	limit := 5
	offset := (page - 1) * limit
	var total int64
	var products []Product

	db.Offset(offset).Limit(limit).Find(&products)

	db.Model(&Product{}).Count(&total)

	//return c.JSON(products)
	return fiber.Map{
		"data": products,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(int(total) / limit)),
		},
	}
}
