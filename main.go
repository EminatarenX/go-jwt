package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"api-auth/api/routes"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Static("/", "./client/dist")

	routes.AuthRouter(app)
	
	app.Listen(":3000")
}