package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		result := map[string]string{"message": "Hello World"}
		return c.JSON(result)
	})

	app.Listen(":8080")
}
