package appserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jdfergason/sonrai/db"
)

func listProducers(c *fiber.Ctx) error {
	producers := db.ListProducers()
	return c.JSON(producers)
}
