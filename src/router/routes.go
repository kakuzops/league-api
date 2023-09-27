package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kakuzops/league-api/src/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/user")

	v1.Get("/", handler.GetAllUsers)
	v1.Get("/:id", handler.GetSingleUser)
	v1.Post("/", handler.CreateUser)
	v1.Put("/:id", handler.UpdateUser)
	v1.Delete("/:id", handler.DeleteUserByID)

	v0 := api.Group("/champion")

	v0.Get("/", handler.GetAllChampions)
	v0.Post("/", handler.CreateChampion)
	v0.Get("/:city", handler.GetChampionsByCity)
}
