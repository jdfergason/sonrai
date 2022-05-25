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
	"encoding/json"
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/penny-vault/sonrai/db"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var createCmd = &cobra.Command{
	Use:   "create [project config]",
	Short: "Create new objects in Sonrai",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, configFn := range args {
			// read project definition
			project := db.Project{}
			body, err := os.ReadFile(configFn)
			if err != nil {
				log.Error().Err(err).Str("FileName", configFn).Msg("could not read project configuration from disk")
				continue
			}
			
			if err := toml.Unmarshal(body, &project); err != nil {
				log.Error().Err(err).Str("FileName", configFn).Msg("could not unmarshal toml")
				continue
			}

			// send to server
			projectJson, err := json.Marshal(project)
			if err != nil {
				log.Error().Err(err).Msg("could not marshal project to JSON")
			}

			cmdUrl := fmt.Sprintf("%s/api/v1/jobs", viper.GetString("api.host"))
			log.Info().Msg("sending project configuration to server")

			// content validation is done on the server
			jobResponse := db.JobResponse{}
			client := resty.New()
			log.Debug().Str("URL", cmdUrl).Msg("HTTP PUT url")
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(projectJson).
				SetResult(&jobResponse).
				Put(cmdUrl)
			if err != nil {
				log.Error().Err(err).Msg("received error from server")
				os.Exit(1)
			}
			if resp.StatusCode() >= 400 {
				log.Error().Int("StatusCode", resp.StatusCode()).Msg("received error from server")
				os.Exit(1)
			}

			log.Info().Str("TraceId", jobResponse.TraceID).Msg("created job")
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
