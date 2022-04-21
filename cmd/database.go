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
package cmd

import (
	"github.com/jdfergason/sonrai/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var dbDriver string
var dbDsn string

// serveCmd represents the serve command
var dbCmd = &cobra.Command{
	Use:   "db-migrate",
	Short: "Initialize or migrate database to latest schema version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().
			Str("database.dsn", viper.GetString("database.dsn")).
			Msg("migrate database to latest schema version")

		db.Migrate(viper.GetString("database.dsn"))
	},
}

func init() {
	cobra.OnInitialize(setupLogging)

	rootCmd.AddCommand(dbCmd)

	dbCmd.PersistentFlags().StringVar(&dbDsn, "database.dsn", "sqlite://sonrai.sqlite3", "database driver to use")
	viper.BindPFlag("database.dsn", dbCmd.PersistentFlags().Lookup("database.dsn"))
}
