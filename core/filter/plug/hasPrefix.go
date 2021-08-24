package plug

import (
	"flybeat/pkg/logging"
	"fmt"
	"strings"
)

func init() {
	Register("HasPrefix", NewHasPrefix)
}

func NewHasPrefix(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 1 {
		if value, ok := data[filed]; ok {
			if strings.HasPrefix(value.(string), render[0]) {
				return value, true
			}
		} else {
			logging.Error(fmt.Sprintf("hasPrefix model:%s not found", filed))
		}
	} else {
		logging.Error("NewHasPrefix len error 1:" + strings.Join(render, ","))
	}
	return nil, false
}
