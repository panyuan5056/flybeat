package plug

import (
	"flybeat/pkg/logging"
	"strings"
)

func init() {
	Register("Convert", NewConvert)
}

func NewConvert(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 1 {
		if row, ok := data[filed]; ok {
			tag := render[0]
			switch tag {
			case "int":
				t, ok := row.(int)
				return t, ok
			case "[]int":
				t, ok := row.([]int)
				return t, ok
			case "string":
				t, ok := row.(string)
				return t, ok
			case "[]string":
				t, ok := row.([]string)
				return t, ok
			case "float32":
				t, ok := row.(float64)
				return t, ok
			case "float64":
				t, ok := row.(float32)
				return t, ok
			default:
				return row, true
			}
		}
	} else {
		logging.Error("NewConvert len error 1:" + strings.Join(render, ","))
	}
	return nil, false
}
