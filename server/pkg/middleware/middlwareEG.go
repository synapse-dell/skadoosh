package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SayHello(c *fiber.Ctx) error {
	fmt.Println("Hello Middleware")
	return c.Next()
}
