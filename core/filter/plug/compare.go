package plug

//gt: greater than 大于
//gte: greater than or equal 大于等于
//lt: less than 小于
//lte: less than or equal 小于等于
//filed:x, compare:gt, value:y
import (
	"flybeat/pkg/logging"
	"strings"
)

func init() {
	Register("Compare", NewCompare)
}

func NewCompare(data map[string]interface{}, filed string, render []string) (interface{}, bool) {
	if len(render) == 2 {
		tag := render[0]
		index2 := render[1]
		if row, ok := data[filed]; ok {
			if row2, ok2 := data[index2]; ok2 {
				value1 := row.(float64)
				value2 := row2.(float64)
				switch tag {
				case ">":
					return nil, value1 > value2
				case ">=":
					return nil, value1 >= value2
				case "<":
					return nil, value1 < value2
				case "<=":
					return nil, value1 <= value2
				case "=":
					return nil, value1 == value2
				}
			}
		} else {
			logging.Error("NewCompare filed not in:" + filed)
		}
	} else {
		logging.Error("NewCompare len error 2:" + strings.Join(render, ","))
	}
	return nil, false
}
