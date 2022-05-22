/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package appserver

import "github.com/gofiber/fiber/v2"

// SetupRoutes setup router api
func setupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/healthz", healthz)

	// Activity
	activity := api.Group("/activity")
	activity.Get("/", healthz)

	// Producers
	producers := api.Group("/producers")
	producers.Get("/", listProducers)
	producers.Put("/", upsertProducer)
	producers.Get("/:slug", healthz)
	producers.Put("/:slug", upsertProducer)
	producers.Delete("/:slug", healthz)

	// Transformers
	transformers := api.Group("/transformers")
	transformers.Get("/", healthz)
	transformers.Put("/", healthz)
	transformers.Get("/:slug", healthz)
	transformers.Put("/:slug", healthz)
	transformers.Delete("/:slug", healthz)

	// Data Records
	records := api.Group("/record-info")
	records.Get("/", healthz)

	// Monitors
	monitors := api.Group("/monitors")
	monitors.Get("/", healthz)
	monitors.Put("/", healthz)
	monitors.Get("/:slug", healthz)
	monitors.Put("/:slug", healthz)
	monitors.Delete("/:slug", healthz)

	// Syncs
	syncs := api.Group("/syncs")
	syncs.Get("/", healthz)
	syncs.Put("/", healthz)
	syncs.Get("/:slug", healthz)
	syncs.Put("/:slug", healthz)
	syncs.Delete("/:slug", healthz)
}
