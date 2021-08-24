package core

import (
	"flybeat/core/topology"
	"flybeat/models"
	"flybeat/pkg/e"
	"flybeat/pkg/patterns"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GrokInfo(c *gin.Context) {
	code := e.SUCCESS
	result := map[string]interface{}{}
	result["rows"] = patterns.Rows
	result["info"] = patterns.Info
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": result,
	})
	c.Abort()
}

func Add(c *gin.Context) {
	var config Config
	code := e.INVALID_PARAMS
	result := map[string]interface{}{}
	if err := c.BindJSON(&config); err == nil {
		code = e.SUCCESS
		models.DB.Create(&config)
		//add chan
		confiChan <- config.tomap()
		result["id"] = config.ID
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": result,
	})
	c.Abort()
}

func TestData(c *gin.Context) {
	var config Config
	code := e.INVALID_PARAMS
	var result []map[string]interface{}
	if err := c.BindJSON(&config); err == nil {
		xconfig := addConfigBox(config.tomap())
		if len(xconfig.Input) == 0 || len(xconfig.Output) == 0 {
			code = e.INVALID_PARAMS
		} else {
			code = e.SUCCESS
			for _, input := range xconfig.Input {
				for i := 0; i < 10; i++ {
					if event, ok := func(input *topology.Input) (map[string]interface{}, bool) {
						select {
						case event := <-input.Messages:
							return event, true
						case <-time.After(time.Second):
							return nil, false
						}
					}(input); ok {
						result = append(result, event)
					} else {
						break
					}
				}
				input.Stop = true
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": result,
	})
	c.Abort()
}

type ConfigFilter struct {
	Data   []map[string]interface{}
	Filter map[string][]map[string]interface{}
}

func TestFilter(c *gin.Context) {
	var config ConfigFilter
	code := e.INVALID_PARAMS
	var result []map[string]interface{}
	if err := c.BindJSON(&config); err == nil {
		code = e.SUCCESS
		filter := topology.GetFilter(config.Filter)
		for _, row := range config.Data {
			if event, ok := filter.Process(row); ok {
				result = append(result, event)
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": result,
	})
	c.Abort()
}

type ConfigMatch struct {
	Data  []map[string]interface{}
	Match map[string]interface{}
}

func TestMatch(c *gin.Context) {
	var config ConfigMatch
	code := e.INVALID_PARAMS
	var result []map[string]interface{}
	if err := c.BindJSON(&config); err == nil {
		code = e.SUCCESS
		box := addConfigBox(map[string]interface{}{"match": config.Match})
		for _, row := range config.Data {
			event := box.Match.Encoder(row["message"].(string))
			result = append(result, event)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
		"result": result,
	})
	c.Abort()
}
