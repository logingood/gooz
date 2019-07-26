package cmd

import (
	"fmt"
	"os"

	"github.com/logingood/gooz/internal/backend/zfile"
	"github.com/logingood/gooz/internal/search"
	"github.com/spf13/afero"
)

func Search(filePath string, field string, pattern string) (results []map[string]interface{}) {
	var appFs = afero.NewOsFs()
	data := zfile.New(appFs, filePath)

	err := data.Open()
	if err != nil {
		fmt.Println("Error while opening file")
		os.Exit(1)
	}

	dataMap, err := data.Read()
	defer data.Close()
	if err != nil {
		fmt.Println("Error while reading file")
		os.Exit(1)
	}

	idx, err := search.BuildIndex(field, dataMap)
	if err != nil {
		fmt.Println("Error while building index")
		os.Exit(1)
	}

	results = search.SearchData(pattern, idx)

	if len(results) < 1 {
		fmt.Println("Search didn't return results, try another word")
		os.Exit(0)
	}

	return results
}
