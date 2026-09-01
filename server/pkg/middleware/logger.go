package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logger(c *fiber.Ctx) error {
	fmt.Println("Request: ", c.Method(), c.Path())
	//c.Method returns the method(GET, POST etc)
	return c.Next()
}
