package routes

import (
	"fiber_sample/handlers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// method GET
	app.Get("/", handlers.ReadData)
	// Method GET by ID
	app.Get("/:id", handlers.ReadDataById)
	// Method POST
	app.Post("/", handlers.InsertData)
	// method DELETE
	app.Delete("/:id", handlers.DeleteData)
	// Method PATCH
	app.Patch("/:id", handlers.PatchData)
}
