package match

import (
	"fmt"
	"regexp"
)

func init() {
	Register("Regex", NewRegexp)
}

func NewRegexp(data string, config map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	match := config["match"].(string)
	fileds := config["fileds"].([]interface{})
	r := regexp.MustCompile(match)
	v := r.FindAllString(data, -1)
	for index, row := range v {
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
