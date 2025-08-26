package main

import (
	"nextts_react_go/internal/repo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*"}))

	app.Get("/", func(c *fiber.Ctx) error {
		//getUser := repo.TB_Import_Error_Main()
		getUser2 := repo.UserMain()

		//result := map[string]string{"message": "Hello World From Go "}
		result := map[string]interface{}{"message": getUser2}
		return c.JSON(result)
	})

	app.Listen(":8080")
}
