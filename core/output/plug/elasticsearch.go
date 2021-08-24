package plug

import (
	"bytes"
	"context"
	"encoding/json"
	"flybeat/core/topology"
	"flybeat/pkg/logging"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esutil"
)

func NewElasticsearchOutput(config map[string]interface{}) (*topology.Output, bool) {
	indexType, ok := config["index_type"]
	if ok {
		indexType = indexType.(string)
	} else {
		indexType = "doc"
	}
	indexHome, ok1 := config["index"]
	indexHome2 := indexHome.(string)
	if !ok1 {
		panic("es index not found")
	} else {
		if strings.Index(indexHome2, "%") > -1 {
			rows := strings.Split(indexHome2, "%")
			if len(rows) == 2 && len(rows[1]) == 0 {
				panic(fmt.Sprintf("index:%s format fail", indexHome2))
			}
		}
	}
	address := strings.Split(config["address"].(string), ",")
	cfg := elasticsearch.Config{
		Addresses: address,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		panic(err.Error())
	}
	output := &topology.Output{
		Messages: make(chan map[string]interface{}, 10),
		Stop:     false,
		Config:   config,
	}
	go func() {
		for {
			indexName := index(indexHome2)
			bi, err := esutil.NewBulkIndexer(esutil.BulkIndexerConfig{
				Index:         indexName,        // The default index name
				Client:        es,               // The Elasticsearch client
				NumWorkers:    runtime.NumCPU(), // The number of worker goroutines
				FlushBytes:    int(5e+6),        // The flush threshold in bytes
				FlushInterval: 30 * time.Second, // The periodic flush interval
			})
			if err != nil {
				logging.Error(err.Error())
			}
			for i := 0; i <= 1000; i++ {
				body := <-output.Messages
				//fmt.Println("body:", body)
				data, err1 := json.Marshal(body)
				if err1 != nil {
					logging.Error(err1.Error())
				}
				err2 := bi.Add(context.Background(), esutil.BulkIndexerItem{
					Action: "index",
					Index:  indexName,
					Body:   bytes.NewReader(data),
				})
				if err2 != nil {
					logging.Error(err2.Error())
				}
			}
			if err3 := bi.Close(context.Background()); err3 != nil {
				logging.Error(err3.Error())
			}
		}
	}()
	return output, true
}

func index(index string) string {
	if strings.Index(index, "%") > -1 {
		rows := strings.Split(index, "%")
		if len(rows) == 2 && len(rows[1]) > 0 {
			return fmt.Sprintf("%s%s", rows[0], time.Now().Format(rows[1]))
		} else {
			return rows[0]
		}
	}
	return index
}
