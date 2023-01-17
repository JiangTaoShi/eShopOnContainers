package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JiangTaoShi/eShopOnContainers/src/ordering/configs"
	"github.com/JiangTaoShi/eShopOnContainers/src/ordering/router"
	"github.com/gin-gonic/gin"
)

func init() {

}

func main() {

	ServerSetting := configs.Get().Server
	gin.SetMode(ServerSetting.RunMode)
	routersInit := router.InitRouter()
	readTimeout := ServerSetting.ReadTimeout
	writeTimeout := ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout * time.Second,
		WriteTimeout:   writeTimeout * time.Second,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("[error] start http server error")
	}
}
