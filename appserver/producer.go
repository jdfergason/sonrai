package appserver

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jdfergason/sonrai/db"
	"github.com/rs/zerolog/log"
)

func listProducers(c *fiber.Ctx) error {
	producers := db.ListProducers()
	return c.JSON(producers)
}

func upsertProducer(c *fiber.Ctx) error {
	var producer db.Producer
	body := c.Request().Body()
	err := json.Unmarshal(body, &producer)
	if err != nil {
		log.Error().Str("Body", string(body)).Msg("cannot unmarshal body")
		return fiber.NewError(405, fmt.Sprintf("cannot unmarshal body: %s", err.Error()))
	}
	return c.JSON(body)
}
