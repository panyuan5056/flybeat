package match

import (
	"encoding/json"
	"flybeat/pkg/logging"
)

func init() {
	Register("Json", NewJson)
}

func NewJson(data string, config map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		logging.Error(err.Error())
	}
	return result
}
