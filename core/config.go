package core

import (
	"flybeat/core/input"
	"flybeat/core/match"
	"flybeat/core/output"
	"flybeat/core/topology"
	"flybeat/pkg/logging"
)

type ConfigBox struct {
	Input  []*topology.Input
	Match  *topology.Match
	Filter *topology.Filter
	Output []*topology.Output
}

func (c *ConfigBox) Beat() {
	for _, input := range c.Input {
		go func(input *topology.Input) {
			for {
				event, _ := input.OnEventBar()
				event["message"] = c.Match.Encoder(event["message"].(string))
				if message, ok := c.Filter.Process(event["message"].(map[string]interface{})); ok {
					for _, output := range c.Output {
						output.Push(message)
					}
				} else {
					logging.Error("log not process:%s", event)
				}
			}
		}(input)
	}
}

func (c *ConfigBox) getMatch(configs map[string]interface{}) {
	for category, config := range configs {
		if handle, ok := match.HandleMatchs[category]; ok {
			c.Match = &topology.Match{Config: config.(map[string]interface{}), Handle: handle}
			break
		}
	}
}

func (c *ConfigBox) getInput(configs map[string]interface{}) {
	for category, config := range configs {
		if xinput, err := input.GetInput(category, config.(map[string]interface{})); err != nil {
			logging.Error(err)
			panic(err)
		} else {
			c.Input = append(c.Input, xinput)
		}
	}

}

func (c *ConfigBox) getFilter(configs map[string][]map[string]interface{}) {
	c.Filter = topology.GetFilter(configs)
}

func (c *ConfigBox) getOutput(configs map[string]interface{}) {
	for category, config := range configs {
		if xoutput, err := output.GetOutput(category, config.(map[string]interface{})); err != nil {
			logging.Error(err)
			panic(err)
		} else {
			c.Output = append(c.Output, xoutput)
		}
	}
}

func addConfigBox(config map[string]interface{}) ConfigBox {
	var c ConfigBox
	if match, ok := config["match"]; ok {
		if ma, ok := match.(map[string]interface{}); ok {
			c.getMatch(ma)
		}
	}
	if input, ok := config["input"]; ok {
		c.getInput(input.(map[string]interface{}))
	}
	if filter, ok := config["filter"]; ok {
		if op, ok := filter.(map[string][]map[string]interface{}); ok {
			c.getFilter(op)
		}
	}
	if output, ok := config["output"]; ok {
		c.getOutput(output.(map[string]interface{}))
	}
	return c
}
