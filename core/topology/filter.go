package topology

import (
	"flybeat/core/filter"
	"flybeat/pkg/logging"
	"fmt"
)

type FilterBox struct {
	fn     func(data map[string]interface{}, config []filter.FilterConfig) (map[string]interface{}, bool)
	config []filter.FilterConfig
}

type Filter struct {
	config interface{}
	bind   []FilterBox
}

func (f *Filter) Process(data map[string]interface{}) (map[string]interface{}, bool) {
	for _, bind := range f.bind {
		if data, ok := bind.fn(data, bind.config); !ok {
			logging.Error(fmt.Sprintf("log:%s fail filter:%s", data, bind.config))
			return nil, false
		}
	}
	return data, true
}

func GetFilter(xconfigs map[string][]map[string]interface{}) *Filter {
	f := &Filter{config: xconfigs}
	for name, configs := range xconfigs {
		if fn, ok := filter.HandleFilter[name]; ok {
			box := FilterBox{fn: fn}
			box.config = filter.LoadFilterConfig(configs)
			f.bind = append(f.bind, box)
		} else {
			err := fmt.Sprintf("module %s not found", fn)
			logging.Error(err)
			panic(err)
		}
	}
	return f
}
