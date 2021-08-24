package filter

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	data := map[string]interface{}{"a": "123", "b": "dfad234"}
	filter := []map[string]interface{}{}
	//filter = append(filter, map[string]interface{}{"filed": "model", "process": "", "value": "Now(2006-01-02)"})
	//filter = append(filter, map[string]interface{}{"filed": "model1", "process": "", "value": "Random(20,10)"})
	filter = append(filter, map[string]interface{}{"filed": "b", "process": "HasSuffix(4)", "value": "Replace(2,q)"})
	config := LoadFilterConfig(filter)
	q, ok := HandleFilter["Modify"](data, config)
	fmt.Println(q)
	fmt.Println(ok)

}
