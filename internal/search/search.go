package search

import (
	"strconv"

	"github.com/logingood/gooz/internal/index"
	log "github.com/sirupsen/logrus"
)

func BuildIndex(key string, data []map[string]interface{}) (idx *index.HashTable, err error) {
	idx = index.New()

	for _, element := range data {
		if element[key] == nil {
			element[key] = ""
		}

		switch element[key].(type) {
		case []interface{}:
			interfaceSlice := element[key].([]interface{})
			// insert multiple keys for tags and etc
			for _, v := range interfaceSlice {
				idx.Insert(v.(string), element)
			}
		case float64:
			err = idx.Insert(strconv.FormatFloat(element[key].(float64), 'f', 0, 64), element)
		case bool:
			err = idx.Insert(strconv.FormatBool(element[key].(bool)), element)
		default:
			err = idx.Insert(element[key], element)
			if err != nil {
				log.Errorf("Error has occured when indexing data %s", err)
			}
		}
	}

	return idx, err
}

func SearchData(key string, h *index.HashTable) (results []map[string]interface{}) {
	values := h.Search(key)
	results = make([]map[string]interface{}, 0)

	if len(values) < 0 {
		log.Errorf("Search didn't return results")
		return results
	}

	for _, val := range values {
		result := val.(map[string]interface{})
		results = append(results, result)
	}

	return results
}
