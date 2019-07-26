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
	"os"
	"strconv"
	"strings"

	"github.com/logingood/gooz/internal/table"
	"github.com/spf13/cobra"
)

var (
	validInput   map[string]bool
	validStrings []string
)

// organizationsCmd represents the organizations command
var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Search organizations table",
	Long: `Allows to search organizations table using any field, including id,
	date, tags, domains and boolean fields. Data is strictly type, error will be
	thrown if incorrect data type is attempted to be passed as an argument to the
	script`,
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
		results := Search("data/organizations.json", args[0], args[1])

		err := table.DrawTable(results)
		if err != nil {
			fmt.Printf("Failed to draw a table - %s\n", err)
			os.Exit(1)
		}

		for _, element := range results {
			users := Search("data/users.json", "organization_id", strconv.FormatFloat(element["_id"].(float64), 'f', 0, 64))
			fmt.Println("WE FOUND USERS FOR THIS ID")
			err := table.DrawTable(users)
			if err != nil {
				fmt.Printf("Failed to draw a table - %s\n", err)
				os.Exit(1)
			}

			fmt.Println("WE FOUND TICKETS FOR THIS ID")
			tickets := Search("data/tickets.json", "organization_id", strconv.FormatFloat(element["_id"].(float64), 'f', 0, 64))
			err = table.DrawTable(tickets)
			if err != nil {
				fmt.Printf("Failed to draw a table - %s\n", err)
				os.Exit(1)
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
