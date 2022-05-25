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
	"fmt"

	"github.com/penny-vault/sonrai/appserver"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var port int

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Sonrai API service",
	Long: `Starts the Sonrai service. The API and user interface run on port 3000 by default but
may be overridden in the [server] section of the config file.

**API** http://localhost:3000/api/v1/<endpoint>
**UI**  http://localhost:3000/ui`,
	Run: func(cmd *cobra.Command, args []string) {
		// start an embedded NATS server
		log.Info().Int("port", viper.GetInt("server.port")).Msg("starting http server")
		app := appserver.Setup()

		listen := fmt.Sprintf(":%d", viper.GetInt("server.port"))
		log.Fatal().Msg(app.Listen(listen).Error())
	},
}

func init() {
	cobra.OnInitialize(setupLogging)

	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntVarP(&port, "port", "p", 3000, "port to run HTTPS server")
	viper.BindPFlag("server.port", serveCmd.Flags().Lookup("port"))
}
