package cmd

import (
	"fmt"
	//	"os"

	"github.com/logingood/gooz/internal/backend/zfile"
	"github.com/logingood/gooz/internal/search"
	"github.com/logingood/gooz/internal/table"
	"github.com/spf13/afero"
)

func searchField(filePath string, field string, pattern string) (results []map[string]interface{}) {
	var appFs = afero.NewOsFs()
	data := zfile.New(appFs, filePath)

	err := data.Open()
	errExit(err, "Error while opening file")

	dataMap, err := data.Read()
	defer data.Close()

	errExit(err, "Error while reading file")

	idx, err := search.BuildIndex(field, dataMap)
	errExit(err, "Error while building index")

	results = search.SearchData(pattern, idx)

	if len(results) < 1 {
		fmt.Printf("Search of %s => %s didn't return results, try another word\n", field, pattern)
		return nil
	}

	return results
}

func drawTable(results []map[string]interface{}) {
	err := table.DrawTable(results)
	errExit(err, "Failed to draw a table")
}

func errExit(err error, message string) {
	if err != nil {
		fmt.Printf("%s - %s\n", message, err)
	}
}
