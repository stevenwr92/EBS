package main

import "github.com/gofiber/fiber/v2"

func main() {
	// Create a new Fiber app
	app := fiber.New()

	// Define a route for the root path ("/") that responds with "Hello, World!"
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Start the Fiber app on port 3000
	app.Listen(":5000")
}
