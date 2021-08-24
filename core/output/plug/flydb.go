package plug

import (
	"bytes"

	"encoding/json"
	"flybeat/core/topology"
	"flybeat/pkg/logging"

	"net/http"
)

func Post(url string, tableName string, data []map[string]interface{}) {
	body := map[string]interface{}{"tableName": tableName, "content": data}
	content, err1 := json.Marshal(body)
	if err1 != nil {
		logging.Error(err1.Error())
	}
	req, err := http.NewRequest("POST", url+"/batchAdd", bytes.NewBuffer(content))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func NewFlydbOutput(config map[string]interface{}) (*topology.Output, bool) {
	tableName := config["index"].(string)
	tableName = index(tableName)
	address := config["address"].(string)

	output := &topology.Output{
		Messages: make(chan map[string]interface{}, 10),
		Stop:     false,
		Config:   config,
	}
	go func(output *topology.Output) {
		for {
			size := len(output.Messages)
			if size > 1000 {
				size = 1000
			}
			tmp := []map[string]interface{}{}
			for i := 0; i <= size; i++ {
				body := <-output.Messages
				tmp = append(tmp, body)
			}
			Post(address, tableName, tmp)
		}
	}(output)
	return output, true
}
