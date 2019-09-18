package main

import (
	"fmt"
	"net/http"
	"test2/pkg/setting"
	"test2/router"
)

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: r,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	_ = s.ListenAndServe()
}
