package routes

import (
	"github.com/gofiber/fiber/v2"
	"api-auth/api/controllers/auth"
	"api-auth/api/middlewares"
)

func AuthRouter(app *fiber.App) {

	client := app.Group("/api/auth")

	client.Post("/register", auth.Register)
	client.Post("/login", auth.Login)
	client.Get("/profile",middlewares.CheckAuth, auth.Profile)
}