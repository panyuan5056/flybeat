package plug

import (
	"flybeat/pkg/logging"
	"strings"
)

func init() {
	Register("Replace", NewReplace)
}

func NewReplace(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 2 {
		if row, ok := data[filed]; ok {
			return strings.Replace(row.(string), render[0], render[1], -1), true
		}
	} else {
		logging.Error("NewReplace len error 2:" + strings.Join(render, ","))
	}
	return nil, false
}
