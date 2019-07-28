// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// organizationsCmd represents the organizations command
var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Search organizations table",
	Long:  `Search by any field from given _id, url, external_id, name, domain_names, created_at, details, shared_tickets, tags`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Please specify query field and search string")
		}

		if validInput[args[0]] == false {
			return fmt.Errorf("Your input is invalid, we only support the following fields: %s\n", strings.Join(validStrings, ", "))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		results := searchField(organizationsFilePath, args[0], args[1])

		drawTable(results)

		for _, element := range results {
			if element["_id"] != nil {
				users := searchField(usersFilePath, "organization_id", strconv.FormatFloat(element["_id"].(float64), 'f', 0, 64))
				drawTable(users)

				tickets := searchField(ticketsFilePath, "organization_id", strconv.FormatFloat(element["_id"].(float64), 'f', 0, 64))
				drawTable(tickets)
			}
		}
	},
}

func init() {
	validInput = make(map[string]bool)

	validStrings = []string{"_id", "url", "external_id", "name", "domain_names", "created_at", "details", "shared_tickets", "tags"}
	for _, v := range validStrings {
		validInput[v] = true
	}

	rootCmd.AddCommand(organizationsCmd)
}
