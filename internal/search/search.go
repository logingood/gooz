package search

import (
	"github.com/logingood/gooz/internal/helpers"
	"github.com/logingood/gooz/internal/index"
	log "github.com/sirupsen/logrus"
)

func BuildIndex(key string, data []map[string]interface{}) (idx *index.HashTable, err error) {
	// Lock mutex to ensure thread safety
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
		default:
			strElemWord, err := helpers.DetectTypeAndStringfy(element[key])
			// Logging this error for information, we still attempt to put this thing in the hash as a key.
			// Hash takes interface{}, so it doesn't really care.
			if err != nil {
				log.Errorf("Error has while converting types, only int, int64, float64, bool, string and []interface{} are supported %s", err)
			}
			err = idx.Insert(strElemWord, element)
			if err != nil {
				log.Errorf("Error has occured when indexing data %s", err)
			}
		}
	}

	return idx, err
}

func SearchData(key string, h *index.HashTable) (results []map[string]interface{}) {
	// Lock mutex to ensure thread safety
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
