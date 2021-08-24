package output

import (
	"errors"
	"flybeat/core/output/plug"
	"flybeat/core/topology"
	"fmt"
	"plugin"
)

func GetOutput(category string, config map[string]interface{}) (*topology.Output, error) {
	var output *topology.Output
	if category == "Elasticsearch" {
		if output, ok := plug.NewElasticsearchOutput(config); ok {
			return output, nil
		}
		return output, errors.New("Elasticsearch init fail")
	} else if category == "Flydb" {
		if output, ok := plug.NewFlydbOutput(config); ok {
			return output, nil
		}
	}
	pluginPath := fmt.Sprintf("./plug/%s.so", category)
	p, err := plugin.Open(pluginPath)
	if err != nil {
		return output, fmt.Errorf("无法打开插件 %s: %v", pluginPath, err)
	}
	newFunc, err := p.Lookup(fmt.Sprintf("New%sOutput", category))
	if err != nil {
		return output, fmt.Errorf("无法找到方法 `New` function in %s: %s", pluginPath, err)
	}
	f, ok := newFunc.(func(map[string]interface{}) (*topology.Output, bool))
	if !ok {
		return output, fmt.Errorf("方法类型失败 %s", pluginPath)
	}
	output, ok = f(config)
	if !ok {
		return output, fmt.Errorf("方法 %s 无法返回接口", pluginPath)
	}
	return output, nil
}
