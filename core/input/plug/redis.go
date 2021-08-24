package plug

import (
	"context"
	"flybeat/core/topology"
	"flybeat/pkg/logging"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
)

func NewRedisInput(config map[string]interface{}) (*topology.Input, bool) {
	r := redis.NewClient(&redis.Options{
		Addr:     config["address"].(string),
		Password: config["password"].(string),
		DB:       cast.ToInt(config["db"]),
	})
	RedisInput := &topology.Input{
		Messages: make(chan map[string]interface{}, 10),
		Stop:     false,
		Config:   config,
	}
	go func() {
		for {
			if RedisInput.Stop {
				fmt.Println("exit curr input")
				break
			}
			if m, err := r.LPop(context.Background(), config["topic"].(string)).Result(); err == nil {
				Meta := make(map[string]interface{})
				Meta["topic"] = time.Now()
				event := map[string]interface{}{}
				event["@metadata"] = map[string]interface{}{"redis": Meta}
				event["message"] = m
				RedisInput.Messages <- event

			} else {
				logging.Error(fmt.Sprintf("redis get data fail sleep 5min%s", err))
				time.Sleep(time.Duration(300000000000))
			}
		}
	}()
	return RedisInput, true
}
