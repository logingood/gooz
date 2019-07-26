package table

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

func DrawTable(inputData []map[string]interface{}) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetRowLine(true)

	if len(inputData) < 1 {
		return fmt.Errorf("Empty input, can't draw a table")
	}

	header, rows := findTheLongestMapAndSliceKeys(inputData)
	table.SetHeader(header)
	table.AppendBulk(rows)

	table.Render()

	fmt.Printf("\n\n")

	return nil
}

func findTheLongestMapAndSliceKeys(inputData []map[string]interface{}) (header []string, rows [][]string) {

	var longestKeys []string
	var keys []string
	var row []string

	rows = make([][]string, 0)

	for _, element := range inputData {

		keys, row = sortMap(element)
		rows = append(rows, row)

		if len(keys) >= len(longestKeys) {
			longestKeys = keys
		}
	}

	return keys, rows
}

func detectTypeAndStringfy(val interface{}) string {
	switch v := val.(type) {
	case string:
		return val.(string)
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', 0, 64)
	case int:
		return strconv.FormatInt(val.(int64), 10)
	case bool:
		return strconv.FormatBool(val.(bool))
	case []interface{}:
		return interfaceToStringSlice(val)
	default:
		log.Errorf("Unknown data type %+v", v)
		return "failed to convert"
	}

}

func interfaceToStringSlice(value interface{}) string {
	interfaceSlice := value.([]interface{})

	strSlice := make([]string, 0)
	for _, element := range interfaceSlice {
		strSlice = append(strSlice, element.(string))
	}

	return strings.Join(strSlice, "\n")
}

func sortMap(unsortedMap map[string]interface{}) (keys []string, rows []string) {
	rows = make([]string, 0)
	for k := range unsortedMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		rows = append(rows, wrapLines(detectTypeAndStringfy(unsortedMap[k])))
	}

	return keys, rows
}

func wrapLines(str string) string {
	strSlice := strings.Split(str, " ")
	str = strings.Join(strSlice, "\n")

	return str
}
