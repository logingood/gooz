package cmd

import (
	"errors"
	"fmt"
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
		results := SearchField(organizationsFilePath, args[0], args[1])

		drawTable(results)
		if searchRelated {
			GetRelatedElements("organizations", results)
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
