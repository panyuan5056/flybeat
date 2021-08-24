package input

import (
	"flybeat/core/input/plug"
	"flybeat/core/topology"
	"fmt"
	"plugin"
)

func GetInput(category string, config map[string]interface{}) (*topology.Input, error) {
	if category == "Redis" {
		if input, ok := plug.NewRedisInput(config); ok {
			return input, nil
		}
	} else if category == "Kafka" {
		if input, ok := plug.NewkafkaInput(config); ok {
			return input, nil
		}
	}
	input := &topology.Input{}
	pluginPath := fmt.Sprintf("./plug/%s", category)
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return input, fmt.Errorf("无法打开插件 %s: %v", pluginPath, err)
	}
	newFunc, err := p.Lookup(fmt.Sprintf("New%sInput", category))
	if err != nil {
		return input, fmt.Errorf("无法找到方法 `New` function in %s: %s", pluginPath, err)
	}
	f, ok := newFunc.(func(map[string]interface{}) (*topology.Input, bool))
	if !ok {
		return input, fmt.Errorf("方法类型失败 %s", pluginPath)
	}
	input, ok = f(config)
	if !ok {
		return input, fmt.Errorf("方法 %s 无法返回接口", pluginPath)
	}
	return input, nil

}
