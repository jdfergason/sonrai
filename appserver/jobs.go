package appserver

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/penny-vault/sonrai/db"
	"github.com/rs/zerolog/log"
)

func listJobs(c *fiber.Ctx) error {
	jobs := db.ListJobs()
	return c.JSON(jobs)
}

func upsertJob(c *fiber.Ctx) error {
	var job db.Job
	body := c.Request().Body()
	err := json.Unmarshal(body, &job)
	if err != nil {
		log.Error().Str("Body", string(body)).Msg("cannot unmarshal body")
		return fiber.NewError(405, fmt.Sprintf("cannot unmarshal body: %s", err.Error()))
	}
	return c.JSON(body)
}
