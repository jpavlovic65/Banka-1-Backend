package main

import (
	"os"

	"banka1.com/db"
	"banka1.com/types"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db.Init()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		response := types.Response{
			Success: true,
			Data:    "Hello, World!",
			Error:   "",
		}
		return c.JSON(response)
	})

	port := os.Getenv("PORT")
	app.Listen("localhost:" + port)
}
