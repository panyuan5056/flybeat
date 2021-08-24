package plug

import (
	"flybeat/pkg/logging"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	Register("Round", NewRound)
}

func NewRound(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 1 {
		if row, ok := data[filed]; ok {
			if t, err := strconv.ParseFloat(fmt.Sprintf("%.%sf", render[0], row), 64); err == nil {
				return t, true
			} else {
				logging.Error(fmt.Sprintf("can't parse float64 %s", row))
			}
		}
	} else {
		logging.Error("NewRound len error 1:" + strings.Join(render, ","))
	}
	return nil, false

}
