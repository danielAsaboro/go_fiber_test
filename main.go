package main

import (
	"fiber_sample/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// first we must setup new Instance of FIber
	app := fiber.New()

	// then we must define our routes function
	routes.Routes(app)

	// create http connection for this api
	app.Listen(":3000")
}
