package cmd

import (
	"errors"
	"fmt"
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
		results := SearchField(ticketsFilePath, args[0], args[1])

		drawTable(results)
		if searchRelated {
			GetRelatedElements("tickets", results)
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
