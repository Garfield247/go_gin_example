package v1

import (
	"github.com/Garfield247/go_gin_example/models"
	"github.com/Garfield247/go_gin_example/pkg/e"
	"github.com/Garfield247/go_gin_example/pkg/logging"
	"github.com/Garfield247/go_gin_example/pkg/setting"
	"github.com/Garfield247/go_gin_example/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != ""{
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(util.GetPage(c),setting.AppSetting.PageSize,maps)
	data["total"]  = models.GetTagTotal(maps)

	util.ResponseWithJson(code,data,c)
}

func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name,"name").Message("名称不能为空")
	valid.MaxSize(name,100,"name").Message("最大长度100字符")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.Range(state,0,1,"state").Message("状态只能为0或1")

	code := e.INVALD_PARAMS
	if ! valid.HasErrors(){
		if ! models.ExistTagByName(name){
			code = e.SUCCESS
			models.AddTag(name,state,createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	} else {
		for _,err := range valid.Errors {
			logging.Fatal("err.KEY:%s;err.MSG:%s;\n",err.Key,err.Message)
		}
	}

	util.ResponseWithJson(code,make(map[string]interface{}),c)
}

func EditTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	var state int = 1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("状态只允许0或1")
	}

	valid.Required(id,"id").Message("ID不能为空")
	valid.Required(modifiedBy,"modieied_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy,100,"modified_by").Message("修改人最大长度100")
	valid.MaxSize(name,100,"name").Message("名称最大长度100")

	code := e.INVALD_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id,data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}else {
		for _,err := range valid.Errors {
			logging.Fatal(err.Key,err.Message)
		}
	}
	util.ResponseWithJson(code,make(map[string]interface{}),c)

}

func DeleteTag(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Required(id,"id").Message("id不能为空")
	valid.Min(id,1,"id").Message("id必须大于1")

	code := e.INVALD_PARAMS

	if ! valid.HasErrors(){
		code = e.SUCCESS
		if models.ExistTagByID(id){
			models.DeleteTag(id)
		}else{
			code = e.ERROR_NOT_EXIST_TAG
		}
	}else {
		for _,err := range valid.Errors {
			logging.Fatal(err.Key,err.Message)
		}
	}
	util.ResponseWithJson(code,make(map[string]interface{}),c)
}
