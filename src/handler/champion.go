package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/kakuzops/league-api/config/database"
	model "github.com/kakuzops/league-api/src/entity"
)

func CreateChampion(c *fiber.Ctx) error {
	db := database.DB.Db

	champion := new(model.Champion)

	err := c.BodyParser(champion)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	err = db.Create(&champion).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "date": err})
	}

	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": champion})
}

func GetAllChampions(c *fiber.Ctx) error {
	db := database.DB.Db
	var champions []model.Champion

	db.Find(&champions)
	if len(champions) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "champions not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "champions Found", "data": champions})
}

func DeleteChampionByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var champion model.Champion
	id := c.Params("id")
	db.Find(&champion, "id = ?", id)
	if champion.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "champion not found", "data": nil})
	}
	err := db.Delete(&champion, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}

func GetSingleChampion(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var champion model.Champion
	db.Find(&champion, "id = ?", id)
	if champion.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "champion not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "champion Found", "data": champion})
}

func GetChampionsByCity(c *fiber.Ctx) error {
	db := database.DB.Db
	city := c.Params("city")
	var champion []model.Champion
	db.Where("city = ?", city).Find(&champion)
	if len(champion) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "champions not found", "data": nil})
	}

	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "City found", "data": city, "Champions": champion})

}
