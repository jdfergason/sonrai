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
package db

import (
	"embed"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "modernc.org/sqlite"

	"github.com/rs/zerolog/log"
)

// Embed a directory
//go:embed migrations/*.sql
var fs embed.FS

func Migrate(dsn string) error {
	d, err := iofs.New(fs, "migrations")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create iofs with database migrations")
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", d, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create migration iofs isntance")
		return err
	}
	currentVersion, dirty, err := m.Version()
	log.Info().Uint("CurrentVersion", currentVersion).Bool("Dirty", dirty).Msg("starting migration")
	err = m.Up()
	if err != nil {
		log.Error().Err(err).Msg("unable to apply database migrations")
		return err
	}
	newVersion, dirty, err := m.Version()
	log.Info().Uint("NewVersion", newVersion).Bool("Dirty", dirty).Msg("completed migration")

	return nil
}
