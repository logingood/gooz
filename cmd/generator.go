package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

// generatorCmd represents the generator command

var genSize int

var generatorCmd = &cobra.Command{
	Use:   "generator",
	Short: "Test data generator",
	Long:  `Generates data for performance analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		appFS := afero.NewOsFs()

		var genOrgs, genUsers, genTickets []map[string]interface{}

		for i := 1; i < genSize; i++ {

			org := genOrg(i)
			genOrgs = append(genOrgs, org)

			user := genUser(i)
			genUsers = append(genUsers, user)

			ticket := genTicket(i)
			genTickets = append(genTickets, ticket)
		}

		genOrgsJson, err := json.Marshal(genOrgs)
		if err != nil {
			fmt.Println(err)
		}

		genUsersJson, err := json.Marshal(genUsers)
		if err != nil {
			fmt.Println(err)
		}

		genTicketsJson, err := json.Marshal(genTickets)
		if err != nil {
			fmt.Println(err)
		}

		afero.WriteFile(appFS, "perfdata/organizations.json", genOrgsJson, 0644)
		afero.WriteFile(appFS, "perfdata/users.json", genUsersJson, 0644)
		afero.WriteFile(appFS, "perfdata/tickets.json", genTicketsJson, 0644)
	},
}

func genOrg(i int) map[string]interface{} {
	org := make(map[string]interface{})
	org["_id"] = i
	org["name"] = fmt.Sprintf("myorg%d", i)
	org["url"] = fmt.Sprintf("https://myorg%d.example.com", i)
	org["external_id"] = fmt.Sprintf("external-id-%d", i)
	org["domain_names"] = []interface{}{fmt.Sprintf("domain1%d", i), fmt.Sprintf("domain2%d", i), fmt.Sprintf("domain3%d", i)}
	org["tags"] = []interface{}{fmt.Sprintf("tag1%d", i), fmt.Sprintf("tag2%d", i), fmt.Sprintf("tag3%d", i)}
	org["created_at"] = fmt.Sprintf("2019-07-30T-%d", i)
	org["details"] = fmt.Sprintf("myorg%d", i)
	org["shared_tickets"] = true

	return org
}

func genUser(i int) map[string]interface{} {
	org := make(map[string]interface{})
	org["_id"] = i
	org["name"] = fmt.Sprintf("myorg%d", i)
	org["url"] = fmt.Sprintf("https://myorg%d.example.com", i)
	org["external_id"] = fmt.Sprintf("external-id-%d", i)
	org["domain_names"] = []interface{}{fmt.Sprintf("domain1%d", i), fmt.Sprintf("domain2%d", i), fmt.Sprintf("domain3%d", i)}
	org["tags"] = []interface{}{fmt.Sprintf("tag1%d", i), fmt.Sprintf("tag2%d", i), fmt.Sprintf("tag3%d", i)}
	org["created_at"] = fmt.Sprintf("2019-07-30T-%d", i)
	org["details"] = fmt.Sprintf("myorg%d", i)
	org["organization_id"] = 10 % i
	org["suspended"] = true
	org["role"] = "agent"

	return org
}

func genTicket(i int) map[string]interface{} {
	ticket := make(map[string]interface{})
	ticket["_id"] = i
	ticket["assignee_id"] = i
	ticket["submitter_id"] = i
	ticket["organization_idd"] = i
	ticket["name"] = fmt.Sprintf("myorg%d", i)
	ticket["url"] = fmt.Sprintf("https://myorg%d.example.com", i)
	ticket["external_id"] = fmt.Sprintf("external-id-%d", i)
	ticket["domain_names"] = []interface{}{fmt.Sprintf("domain1%d", i), fmt.Sprintf("domain2%d", i), fmt.Sprintf("domain3%d", i)}
	ticket["tags"] = []interface{}{fmt.Sprintf("tag1%d", i), fmt.Sprintf("tag2%d", i), fmt.Sprintf("tag3%d", i)}
	ticket["created_at"] = fmt.Sprintf("2019-07-30T-%d", i)
	ticket["details"] = fmt.Sprintf("myorg%d", i)
	ticket["shared_tickets"] = true

	return ticket
}

func init() {
	rootCmd.AddCommand(generatorCmd)
	generatorCmd.PersistentFlags().IntVar(&genSize, "size", 100000, "generate 100k records by default or whatever you give")
}
