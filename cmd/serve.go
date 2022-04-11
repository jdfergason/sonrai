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
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Sonrai API service",
	Long: `Starts the Sonrai service. The API and user interface run on port 3000 by default but
may be overridden in the [server] section of the config file.

**API** http://localhost:3000/api/v1/<endpoint>
**UI**  http://localhost:3000/ui`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")

		// start an embedded NATS server
		opts := &server.Options{}
		ns, err := server.NewServer(opts)

		if err != nil {
			panic(err)
		}

		go ns.Start()

		if !ns.ReadyForConnections(4 * time.Second) {
			panic("not ready for connection")
		}

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
