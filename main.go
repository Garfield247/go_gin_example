package main

import (
	"fmt"
	"github.com/Garfield247/go_gin_example/pkg/setting"
	"github.com/Garfield247/go_gin_example/routers"
	//"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



func main() {

	//router := gin.Default()
	//router.GET("/test", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"msg":"OK!",
	//	})
	//})
	router := routers.InitRouter()
	log.Printf(":%d",setting.ServerSetting.HttpPort)
	s := &http.Server{
		Addr: fmt.Sprintf(":%d",setting.ServerSetting.HttpPort),
		Handler: router,
		ReadTimeout: setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1<<20,
	}
	s.ListenAndServe()
}