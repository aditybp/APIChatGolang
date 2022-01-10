package main

import (
	"APICHATGOLANG/database"
	"APICHATGOLANG/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	database.ConnectionDB()
	routes.Setup(app)
	app.Listen(":8000")
}
