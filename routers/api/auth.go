package api

import (
	"github.com/Garfield247/go_gin_example/models"
	"github.com/Garfield247/go_gin_example/pkg/e"
	"github.com/Garfield247/go_gin_example/pkg/logging"
	"github.com/Garfield247/go_gin_example/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username,Password: password}
	ok,_ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALD_PARAMS
	if ok {
		isExist := models.CheckAuth(username,password)
		if isExist {
			token, err := util.GenerateToken(username,password)
			if err != nil {

			} else {
				data["token"] = token
				code = e.SUCCESS

			}
		} else {
			code = e.ERROR_AUTH
		}
		}else {
			for _, err := range valid.Errors {
				logging.Debug(err.Key,err.Message)
			}
	}
	util.ResponseWithJson(code,data,c)
}
