package plug

import (
	"context"
	"flybeat/core/topology"
	"flybeat/pkg/logging"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func NewkafkaInput(config map[string]interface{}) (*topology.Input, bool) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   config["brokers"].([]string),
		Topic:     config["topic"].(string),
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	KafkaInput := &topology.Input{
		Messages: make(chan map[string]interface{}, 10),
		Stop:     false,
		Config:   config,
	}
	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				logging.Error(fmt.Sprintf("kafka获取数据失败%s", err))
				break
			}
			Meta := make(map[string]interface{})
			Meta["topic"] = m.Topic
			Meta["partition"] = m.Partition
			Meta["offset"] = m.Offset
			Meta["topic"] = time.Now()
			event := map[string]interface{}{}
			event["@metadata"] = map[string]interface{}{"kafka": Meta}
			event["message"] = m.Value
			KafkaInput.Messages <- event
		}
		if err := r.Close(); err != nil {
			logging.Error(fmt.Sprintf("关闭失败:%s", err))

		}
	}()
	return KafkaInput, true
}
