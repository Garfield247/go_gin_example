package util

import (
	"github.com/Garfield247/go_gin_example/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int
	Msg string
	Data interface{}
}


func ResponseWithJson(code int,data interface{},c *gin.Context) {
	c.JSON(http.StatusOK,&Response{
		Code: code,
		Msg: e.GetMsg(code),
		Data: data,
	})
}