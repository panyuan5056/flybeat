package filter

import (
	"flybeat/core/filter/plug"
	"fmt"
	"strings"
)

//{"filed":"model", "process":"hasPrefix(mode, a)", "value":"date(x, %h:%s:%m)"}

var HandleFilter map[string]func(data map[string]interface{}, renders []FilterConfig) (map[string]interface{}, bool) = make(map[string]func(data map[string]interface{}, renders []FilterConfig) (map[string]interface{}, bool))

func Register(category string, fn func(data map[string]interface{}, renders []FilterConfig) (map[string]interface{}, bool)) {
	HandleFilter[category] = fn
}

type Render struct {
	fn      func(data map[string]interface{}, filed string, render []string) (interface{}, bool)
	value   interface{}
	args    []string
	process bool
}

func loadRender(config interface{}) Render {
	render := Render{process: true}
	if val, ok := config.(string); ok {
		if strings.Contains(val, "(") {
			rows := strings.Split(strings.Replace(val, ")", "", -1), "(")
			if xfn, ok := plug.HandleRender[rows[0]]; ok {
				render.fn = xfn
				render.args = strings.Split(rows[1], ",")
				render.process = true
			} else {
				panic(fmt.Sprintf("fn:%s not in fliter plug", rows[0]))
			}

		} else {
			render.value = config
			render.process = false
		}
	} else {
		render.value = config
		render.process = false
	}
	return render
}

type FilterConfig struct {
	Filed   string
	Process Render
	Value   Render
}

func (f *FilterConfig) status(data map[string]interface{}) bool {
	if f.Process.process {
		if _, ok := f.Process.fn(data, f.Filed, f.Process.args); ok {
			return true
		}
	} else {
		return true
	}
	return false
}

func (f *FilterConfig) value(data map[string]interface{}) (interface{}, bool) {
	if f.Value.process {
		return f.Value.fn(data, f.Filed, f.Value.args)
	}
	return f.Value.value, true
}

func LoadFilterConfig(configs []map[string]interface{}) []FilterConfig {
	result := []FilterConfig{}
	for _, config := range configs {
		f := FilterConfig{}
		if filed, ok := config["filed"]; ok {
			f.Filed = filed.(string)
		} else {
			panic(fmt.Sprintf("config:%s filed must fill", config))
		}
		if value, ok := config["value"]; ok {
			f.Value = loadRender(value)
		} else {
			panic(fmt.Sprintf("config:%s process must fill", config))
		}
		if process, ok := config["process"]; ok {
			f.Process = loadRender(process)
		} else {
			f.Process = Render{process: false}
		}
		result = append(result, f)
	}
	return result
}
