package match

import (
	"fmt"
	"strings"
)

func init() {
	Register("Split", NewSplit)
}

func NewSplit(data string, config map[string]interface{}) map[string]interface{} {
	seq := config["seq"].(string)
	fileds := config["fileds"].([]interface{})
	result := map[string]interface{}{}
	for index, row := range strings.Split(data, seq) {
		tag := ""
		if index >= len(fileds) {
			tag = fmt.Sprintf("custom_%d", index)
		} else {
			tag = fileds[index].(string)
		}
		result[tag] = row
	}
	return result
}
