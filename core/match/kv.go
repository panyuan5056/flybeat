package match

import (
	"flybeat/pkg/logging"
	"strings"
)

func init() {
	Register("Kv", NewKv)
}

func NewKv(data string, config map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	seq := config["seq"].(string)
	delimiter := config["delimiter"].(string)
	for _, row := range strings.Split(data, delimiter) {
		item := strings.Split(row, seq)
		if len(item) == 2 {
			result[item[0]] = item[1]
		} else {
			logging.Error("解析报错")
		}
	}
	return result
}
