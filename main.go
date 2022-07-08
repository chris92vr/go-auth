package main

import (
	"github.com/chris92vr/go-auth/database"
	"github.com/chris92vr/go-auth/routes"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen("localhost:8080")
}
