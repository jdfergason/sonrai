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
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var producerCreateCmd = &cobra.Command{
	Use:   "producer [name] [cmd] [args]",
	Short: "Create a new producer",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		slug := slug.Make(name)
		cmdExe := args[1]
		cmdArgs := []string{}
		if len(args) > 2 {
			cmdArgs = args[2:]
		}

		log.Info().Str("Name", args[0]).Str("Slug", slug).Str("Command", cmdExe).Strs("Args", cmdArgs).Msg("creating producer")
		/*
			client := resty.New()
			cmdUrl := fmt.Sprintf("%s/api/v1/producers", viper.GetString("api.host"))
			log.Debug().Str("URL", cmdUrl).Msg("fetching producers from API")
			resp, err := client.R().Get(cmdUrl)
			if err != nil {
				log.Error().Str("OriginalError", err.Error()).Msg("received error from server")
				os.Exit(1)
			}
			if resp.StatusCode() >= 400 {
				log.Error().Int("StatusCode", resp.StatusCode()).Msg("received error from server")
				os.Exit(1)
			}

			body := resp.Body()
			producers := []*db.Producer{}
			err = json.Unmarshal(body, &producers)
			if err != nil {
				log.Error().Str("OriginalError", err.Error()).Msg("failed de-serializing JSON")
				os.Exit(1)
			}

			for _, producer := range producers {
				fmt.Printf("%+v\n", producer)
			}
		*/
	},
}

func init() {
	createCmd.AddCommand(producerCreateCmd)
}
