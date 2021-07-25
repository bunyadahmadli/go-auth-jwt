package main

import (
	"auth-go/database"
	"auth-go/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)



func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		}))

	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}
