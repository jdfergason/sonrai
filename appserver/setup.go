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

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var fiberApp *fiber.App

// Embed a directory
//go:embed dist/*
var embedDirUi embed.FS

func configureCors(app *fiber.App) {
	// Configure CORS
	corsConfig := cors.Config{
		AllowOrigins: "http://localhost:8080, https://www.pennyvault.com, https://beta.pennyvault.com",
		AllowHeaders: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}
	app.Use(cors.New(corsConfig))
}

func Setup() *fiber.App {
	// start the fiber server
	fiberApp = fiber.New()
	configureCors(fiberApp)

	setupRoutes(fiberApp)

	fiberApp.Use("/", filesystem.New(filesystem.Config{
		Root:       http.FS(embedDirUi),
		PathPrefix: "dist",
		Browse:     true,
	}))

	return fiberApp
}
