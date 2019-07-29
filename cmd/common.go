package cmd

import (
	"fmt"

	"github.com/logingood/gooz/internal/backend/zfile"
	"github.com/logingood/gooz/internal/helpers"
	"github.com/logingood/gooz/internal/search"
	"github.com/logingood/gooz/internal/table"
	"github.com/spf13/afero"
)

func SearchField(filePath string, field string, pattern string) (results []map[string]interface{}) {
	var appFs = afero.NewOsFs()
	data := zfile.New(appFs, filePath)

	err := data.Open()
	printErr(err, "Error while opening file")

	dataMap, err := data.Read()
	defer data.Close()

	printErr(err, "Error while reading file")

	idx, err := search.BuildIndex(field, dataMap)
	printErr(err, "Error while building index")

	results = search.SearchData(pattern, idx)

	if len(results) < 1 {
		fmt.Printf("Search of %s => %s didn't return results, try another word\n", field, pattern)
		return nil
	}

	return results
}

func GetRelatedElements(searchGroup string, mapToSearch []map[string]interface{}) {
	for _, element := range mapToSearch {
		switch searchGroup {
		case "organizations":
			if element["_id"] != nil {
				elementStr, err := helpers.DetectTypeAndStringfy(element["_id"])
				if err == nil {
					users := SearchField(usersFilePath, "organization_id", elementStr)
					drawTable(users)

					tickets := SearchField(ticketsFilePath, "organization_id", elementStr)
					drawTable(tickets)
				}
			}
		case "users":
			if element["organization_id"] != nil {
				elementStr, err := helpers.DetectTypeAndStringfy(element["organization_id"])
				if err == nil {
					orgs := SearchField(organizationsFilePath, "_id", elementStr)
					drawTable(orgs)
				}
			}

			if element["_id"] != nil {
				elementStr, err := helpers.DetectTypeAndStringfy(element["_id"])
				if err == nil {
					tickets := SearchField(ticketsFilePath, "assignee_id", elementStr)
					drawTable(tickets)
					tickets = SearchField(ticketsFilePath, "submitter_id", elementStr)
					drawTable(tickets)
				}
			}
		case "tickets":
			if element["organization_id"] != nil {
				elementStr, err := helpers.DetectTypeAndStringfy(element["organization_id"])
				if err == nil {
					orgs := SearchField(organizationsFilePath, "_id", elementStr)
					drawTable(orgs)
				}
			}

			if element["assignee_id"] != nil {
				elementStr, err := helpers.DetectTypeAndStringfy(element["assignee_id"])
				if err == nil {
					users := SearchField(usersFilePath, "_id", elementStr)
					drawTable(users)
				}
			}

			if element["submitter_id"] != nil {
				elementStr, err := helpers.DetectTypeAndStringfy(element["submitter_id"])
				if err == nil {
					users := SearchField(usersFilePath, "_id", elementStr)
					drawTable(users)
				}
			}
		}
	}
}

func drawTable(results []map[string]interface{}) {
	err := table.DrawTable(results)
	printErr(err, "Failed to draw a table")
}

func printErr(err error, message string) {
	if err != nil {
		fmt.Printf("%s - %s\n", message, err)
	}
}
