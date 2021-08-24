package plug

import (
	"flybeat/pkg/logging"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func init() {
	Register("Random", NewRandom)
}

func NewRandom(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 2 {
		rand.Seed(time.Now().UTC().UnixNano())
		if max, err1 := strconv.Atoi(render[0]); err1 != nil {
			logging.Error(fmt.Sprintf("parse int fail:%s", render[0]))
		} else {
			if min, err2 := strconv.Atoi(render[1]); err2 != nil {
				logging.Error(fmt.Sprintf("parse int fail:%s", render[1]))
			} else {
				fmt.Println(max, min, max-min)
				return min + rand.Intn(max-min), true
			}
		}
	} else {
		logging.Error("NewRandom len error 2:" + strings.Join(render, ","))
	}
	return nil, false
}
