package routes

import (
	"auth-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App)  {
	api := app.Group("/api")
	api.Post("/register",controllers.Register)
	api.Post("/login",controllers.Login)
	api.Get("/user",controllers.User)

}
