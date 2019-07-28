package cmd

import (
	"errors"
	"fmt"
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
		if searchRelated {
			getRelatedElements("users", results)
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
