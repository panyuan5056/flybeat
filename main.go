package main

import (
	"flybeat/core"
	"flybeat/pkg/logging"
	"flybeat/pkg/setting"
	"flybeat/routers"
	"fmt"
	"net/http"
)

func main() {
	go core.Run()
	logging.Info("服务开始")
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: setting.MaxHeaderBytes,
	}
	s.ListenAndServe()
}
