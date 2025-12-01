package routes

import (
	"todo/src/controllers"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	auth := app.Group("/auth")
	auth.Post("/register", controllers.RegisterUser) //auth/register
	auth.Post("/login", controllers.LoginUser) //auth/login
	auth.Post("/logout", controllers.LogoutUser) //auth/logout
}