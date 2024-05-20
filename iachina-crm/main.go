package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"xlsqac/iachina-crm/pkg/setting"
	"xlsqac/iachina-crm/routers"
)

func main() {
	r := routers.InitRouter()
	gin.SetMode(gin.DebugMode)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
