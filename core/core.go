package core

import (
	"flybeat/pkg/logging"
)

var confiChan = make(chan map[string]interface{}, 3)

func Run() {
	logging.Info("beat start")
	if configs := loadConfig(); len(configs) > 0 {
		for _, config := range configs {
			confiChan <- config
		}
	} else {
		logging.Error("setting nil")
	}

	for config := range confiChan {
		go func(config map[string]interface{}) {
			//get output
			xconfig := addConfigBox(config)
			xconfig.Beat()
		}(config)
	}
}
