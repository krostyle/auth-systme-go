package main

import (
	"github.com/gofiber/fiber/v2"
	setup "github.com/krostyle/auth-systme-go/src/infrastructure/config"
)

func main() {
	app := fiber.New()
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
	setup.Setup(app)
}
