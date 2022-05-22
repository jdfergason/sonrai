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

	"github.com/go-resty/resty/v2"
	"github.com/gosimple/slug"
	"github.com/jdfergason/sonrai/db"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var producerCreateCmd = &cobra.Command{
	Use:   "producer [name] [cmd] [args]",
	Short: "Create a new producer",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		producer := db.Producer{
			Name:      args[0],
			Slug:      slug.Make(args[0]),
			Command:   args[1],
			Arguments: args[2:],
		}

		producerResponse := db.Producer{}

		producerJson, err := json.Marshal(producer)
		if err != nil {
			log.Error().Err(err).Msg("could not marshal producer to JSON")
		}

		log.Info().
			Str("Name", producer.Name).
			Str("Slug", producer.Slug).
			Str("Command", producer.Name).
			Strs("Args", producer.Arguments).
			Msg("creating producer")

		client := resty.New()
		cmdUrl := fmt.Sprintf("%s/api/v1/producers", viper.GetString("api.host"))
		log.Debug().Str("URL", cmdUrl).Msg("put to producer API")
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(producerJson).
			SetResult(&producerResponse).
			Put(cmdUrl)
		if err != nil {
			log.Error().Err(err).Msg("received error from server")
			os.Exit(1)
		}
		if resp.StatusCode() >= 400 {
			log.Error().Int("StatusCode", resp.StatusCode()).Msg("received error from server")
			os.Exit(1)
		}

		log.Info().Str("ID", producer.ID.String()).Msg("created producer")
	},
}

func init() {
	createCmd.AddCommand(producerCreateCmd)
}
