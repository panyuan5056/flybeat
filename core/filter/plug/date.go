package plug

import (
	"flybeat/pkg/logging"
	"strings"
	"time"
)

func init() {
	Register("Date", NewDate)
	Register("Now", NewNow)
}

func NewDate(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 1 {
		if row, ok := data[filed]; ok {
			if t, err := time.Parse(row.(string), render[1]); err == nil {
				return t, true
			}
		}
	} else {
		logging.Error("NewDate len error 1:" + strings.Join(render, ","))
	}
	return nil, false
}

func NewNow(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 1 {
		return time.Now().Format(render[0]), true
	} else {
		return time.Now(), true
	}
	return nil, false
}
