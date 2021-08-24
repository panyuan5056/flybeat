package core

import (
	"encoding/json"
	"flybeat/models"
	"flybeat/pkg/logging"
	"time"
)

type Config struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Match     string    `json:"match"`
	Input     string    `json:"input"  binding:"required"`
	Filter    string    `json:"filter" `
	Output    string    `json:"output"  binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Status    int       `json:"status" `
	Name      string    `json:"name"`
}

func (c *Config) tomap() map[string]interface{} {
	return map[string]interface{}{
		"match":  c.tojson(c.Match),
		"input":  c.tojson(c.Input),
		"filter": c.tofilter(c.Filter),
		"output": c.tojson(c.Output),
	}
}

func (c *Config) tojson(data string) map[string]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		logging.Error(err.Error())
		panic(err.Error())
	}
	return result
}

func (c *Config) tofilter(data string) map[string][]map[string]interface{} {
	var result map[string][]map[string]interface{}
	if err := json.Unmarshal([]byte(data), &result); err != nil {
		logging.Error(err.Error())
		panic(err.Error())
	}
	return result
}

func loadConfig() []map[string]interface{} {
	var configs []Config
	var results []map[string]interface{}
	models.DB.Where("status = ?", 1).Order("created_at desc").Find(&configs)
	for _, config := range configs {
		tmp := map[string]interface{}{}
		tmp["id"] = config.ID
		tmp["filter"] = config.tofilter(config.Filter)
		tmp["output"] = config.tojson(config.Output)
		tmp["input"] = config.tojson(config.Input)
		tmp["match"] = config.tojson(config.Match)
		results = append(results, tmp)
	}
	return results
}

func init() {
	models.DB.AutoMigrate(&Config{})
}
