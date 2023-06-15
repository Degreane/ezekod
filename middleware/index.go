package middleware

import "github.com/gofiber/fiber/v2"

var (
	MiddleWares map[string]func(c *fiber.Ctx) error = make(map[string]func(c *fiber.Ctx) error)
)

func getDefault(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"path":   c.Path(),
		"method": c.Method(),
	})
}
func postDefault(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"path":   c.Path(),
		"method": c.Method(),
	})
}

func Init() {
	MiddleWares["getDefault"] = getDefault
	MiddleWares["postDefault"] = postDefault
}
