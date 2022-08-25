package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func Home() fiber.Handler {

	return func(c *fiber.Ctx) error {

		c.Status(http.StatusOK)
		
		return c.JSON(map[string]any{
			"msg": "ok",
		})
	}

}
