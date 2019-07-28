package helpers

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Type detector, returns a parsed string if type detected
func DetectTypeAndStringfy(val interface{}) (strConverted string, err error) {
	switch v := val.(type) {
	case string:
		return val.(string), nil
	case float64:
		return strconv.FormatFloat(val.(float64), 'f', 0, 64), nil
	case int:
		return strconv.FormatInt(int64(val.(int)), 10), nil
	case int64:
		return strconv.FormatInt(val.(int64), 10), nil
	case bool:
		return strconv.FormatBool(val.(bool)), nil
	case []interface{}:
		return interfaceToStringSlice(val), nil
	default:
		log.Errorf("Unknown data type %+v", v)
		return "failed to convert", fmt.Errorf("Failed to detect type")
	}
}

// Detect the longest map, because values could have amount of keys
// return to slices []string and [][]string which will be used as header and rows for the table
func FindTheLongestMapAndSliceKeys(inputData []map[string]interface{}) (header []string, rows [][]string) {

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

	return longestKeys, rows
}

// Maps are unsorted in Go, we need to sort them
func sortMap(unsortedMap map[string]interface{}) (keys []string, rows []string) {
	rows = make([]string, 0)
	for k := range unsortedMap {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, k := range keys {
		line, err := DetectTypeAndStringfy(unsortedMap[k])
		if err != nil {
			log.Errorf("Could detect type, %s", err)
		}

		rows = append(rows, wrapLines(line))
	}

	return keys, rows
}

// To make tables look pretty
func wrapLines(str string) string {
	strSlice := strings.Split(str, " ")
	str = strings.Join(strSlice, "\n")

	return str
}

// Prepare a slice of strings
func interfaceToStringSlice(value interface{}) string {
	interfaceSlice := value.([]interface{})

	strSlice := make([]string, 0)
	for _, element := range interfaceSlice {
		strSlice = append(strSlice, element.(string))
	}

	return strings.Join(strSlice, "\n")
}
