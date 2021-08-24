package match

import (
	"fmt"
	"testing"
)

func TestMatch(t *testing.T) {
	data := "{\"a\":123}"
	config := map[string]interface{}{"seq": "|", "fileds": []interface{}{"index", "category", "date", "msg", "id"}}
	q := HandleMatchs["Json"](data, config)
	fmt.Println(q)

	data1 := `127.0.0.1 - - [23/Apr/2014:22:58:32 +0200] "GET /index.php HTTP/1.1" 404 207`
	config1 := map[string]interface{}{"match": "%{COMMONAPACHELOG}"}
	q1 := HandleMatchs["Grok"](data1, config1)
	for k, v := range q1 {
		fmt.Println("k:", k)
		fmt.Println("v:", v)
	}
}
