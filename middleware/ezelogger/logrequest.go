package ezelogger

import "github.com/gofiber/fiber/v2"

func LogRequest(c *fiber.Ctx) error {
	Ezelogger.Printf("% +v", c)
	return c.Next()
}
