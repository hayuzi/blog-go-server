package main

import (
	"fmt"
	"net/http"

	"blog-go-server/pkg/setting"
	"blog-go-server/routers"
	"log"
	"blog-go-server/models"
)

func main() {

	// 初始化设置
	setting.Setup()
	models.Setup()

	// 提取基础配置
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	// 开启服务
	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	s.ListenAndServe()
}
