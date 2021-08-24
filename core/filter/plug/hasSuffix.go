package plug

import (
	"flybeat/pkg/logging"

	"strings"
)

func init() {
	Register("HasSuffix", NewHasSuffix)
}

func NewHasSuffix(data map[string]interface{}, filed string, render []string) (interface{}, bool) {

	if len(render) == 1 {
		if value, ok := data[filed]; ok {
			if strings.HasSuffix(value.(string), render[0]) {
				return value, true
			}
		}
	} else {
		logging.Error("NewHasSuffix len error 1:" + strings.Join(render, ","))
	}
	return nil, false
}
