package match

import (
	"flybeat/pkg/logging"

	"github.com/vjeantet/grok"
)

func init() {
	Register("Grok", NewGrok)
}

func NewGrok(data string, config map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	parse := ""
	if match, ok := config["match"]; ok && len(match.(string)) > 0 {
		parse = match.(string)
	} else {
		parse = config["rule"].(string)
	}
	g, _ := grok.New()
	values, err := g.Parse(parse, data)
	if err == nil {
		for k, v := range values {
			result[k] = v
		}
	} else {
		logging.Error(err)
	}
	return result
}
