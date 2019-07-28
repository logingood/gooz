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

// ticketsCmd represents the tickets command
var ticketsCmd = &cobra.Command{
	Use:   "tickets",
	Short: "Tickets search",
	Long:  `Long desc goes here`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Please specify query field and search string")
		}

		if validInputTickets[args[0]] == false {
			return fmt.Errorf("Your input is invalid, we only support the following fields: %s\n", strings.Join(validStringsTickets, ", "))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		results := searchField("data/tickets.json", args[0], args[1])

		drawTable(results)

		for _, element := range results {
			orgs := searchField("data/organizations.json", "_id", strconv.FormatFloat(element["organization_id"].(float64), 'f', 0, 64))
			drawTable(orgs)

			users := searchField("data/users.json", "_id", strconv.FormatFloat(element["assignee_id"].(float64), 'f', 0, 64))
			drawTable(users)

			users = searchField("data/users.json", "_id", strconv.FormatFloat(element["submitter_id"].(float64), 'f', 0, 64))
			drawTable(users)
		}
	},
}

func init() {
	validInputTickets = make(map[string]bool)

	validStringsTickets = []string{"_id", "url", "external_id", "created_at", "type", "subject", "description", "priority", "status", "submitter_id", "assignee_id", "organization_id", "tags", "has_incidents", "due_at", "via"}
	for _, v := range validStringsTickets {
		validInputTickets[v] = true
	}

	rootCmd.AddCommand(ticketsCmd)
}
