package table

import (
	"fmt"
	"os"

	"github.com/logingood/gooz/internal/helpers"
	"github.com/olekukonko/tablewriter"
)

func DrawTable(inputData []map[string]interface{}) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	if len(inputData) < 1 {
		return fmt.Errorf("Empty input, can't draw a table")
	}

	header, rows := helpers.FindTheLongestMapAndSliceKeys(inputData)
	table.SetHeader(header)
	table.AppendBulk(rows)

	table.Render()

	fmt.Printf("\n\n")

	return nil
}
