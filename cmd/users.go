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

// usersCmd represents the users command
var usersCmd = &cobra.Command{
	Use:   "users",
	Short: "search users table",
	Long:  `blah`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("Please specify query field and search string")
		}

		if validInputUsers[args[0]] == false {
			return fmt.Errorf("Your input is invalid, we only support the following fields: %s\n", strings.Join(validStringsUsers, ", "))
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		results := searchField(usersFilePath, args[0], args[1])

		drawTable(results)

		for _, element := range results {
			if element["organization_id"] != nil {
				orgs := searchField(organizationsFilePath, "_id", strconv.FormatFloat(element["organization_id"].(float64), 'f', 0, 64))
				drawTable(orgs)
			}

			if element["_id"] != nil {
				tickets := searchField(ticketsFilePath, "assignee_id", strconv.FormatFloat(element["_id"].(float64), 'f', 0, 64))
				drawTable(tickets)
				tickets = searchField(ticketsFilePath, "submitter_id", strconv.FormatFloat(element["_id"].(float64), 'f', 0, 64))
				drawTable(tickets)
			}
		}
	},
}

func init() {
	validInputUsers = make(map[string]bool)

	validStringsUsers = []string{"_id", "url", "external_id", "name", "alias", "created_at", "active", "verified", "shared", "locale", "timezone", "last_login_at", "email", "phone", "signature", "organization_id", "tags", "suspended", "role"}
	for _, v := range validStringsUsers {
		validInputUsers[v] = true
	}

	rootCmd.AddCommand(usersCmd)
}
