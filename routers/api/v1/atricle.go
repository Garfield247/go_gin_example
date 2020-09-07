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

func GetArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id,1,"id").Message("ID必须大于0")

	code := e.INVALD_PARAMS
	var data interface{}

	if ! valid.HasErrors(){

		if models.ExistArticleByID(id){
			data = models.GetArticle(id)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}

	}
	util.ResponseWithJson(code,data,c)
}

func GetArticles(c *gin.Context)  {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state,0,1,"state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId,1,"tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALD_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		data["lists"] = models.GetArticles(util.GetPage(c),setting.AppSetting.PageSize,maps)
		data["total"] = models.GetArticleTotal(maps)
	}else {
		for _, err :=range valid.Errors {
			logging.Fatal(err.Key,err.Message)
		}
	}
	util.ResponseWithJson(code,data,c)
}

func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	context := c.Query("context")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("stste","0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId,1,"tag_id").Message("标签ID必须大于0")
	valid.Required(title,"title").Message("文章标题不能为空")
	valid.MaxSize(title,100,"title").Message("标题不能超100字符")
	valid.Required(desc,"desc").Message("文章描述不能为空")
	valid.MaxSize(desc,255,"desc").Message("描述不能超255字符")
	valid.Required(context,"context").Message("文章正文不能为空")
	valid.Required(createdBy,"created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy,100,"created_by").Message("创建人名称不能超过100字符")
	valid.Range(state,0,1,"state").Message("状态只能是0或1")

	code := e.INVALD_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["context"] = context
			data["created_by"] = createdBy
			data["state"] = state

			models.AddArticle(data)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	} else {
		for _,err := range valid.Errors {
			logging.Fatal(err.Key,err.Message)
		}
	}
	util.ResponseWithJson(code,make(map[string]interface{}),c)

}

func EditArticle(c *gin.Context)  {
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	context := c.Query("context")
	modifiedBy := c.Query("modified_by")

	var state int = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		valid.Range(state,0,1,"state").Message("状态只能取0或1")
	}
	valid.Min(id,1,"id").Message("ID必须大于1")
	valid.MaxSize(title,100,"title").Message("标题最长为100字符")
	valid.MaxSize(desc,255,"desc").Message("描述不能超255字符")
	valid.MaxSize(context,65535,"context").Message("文章正文不能超过65535字符")
	valid.Required(modifiedBy,"modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy,100,"modified_by").Message("修改人名称不能超过100字符")

	code := e.INVALD_PARAMS
	if ! valid.HasErrors() {
		if models.ExistArticleByID(id){
			if models.ExistTagByID(tagId){
				data := make(map[string]interface{})
				if tagId > 0 {
					data["tag_id"] = tagId
				}
				if title != ""{
					data["title"] = title
				}
				if desc != ""{
					data["desc"] = desc
				}
				if context != "" {
					data["context"] = context
				}

				data["modified_by"] = modifiedBy

				models.EditArticle(id,data)
				code = e.SUCCESS
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else  {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}else {
		for _,err := range valid.Errors {
			logging.Fatal(err.Key,err.Message)
		}
	}
	util.ResponseWithJson(code,make(map[string]interface{}),c)

}

func DeleteArticle(c *gin.Context)  {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id,1,"id").Message("ID不能小于1")

	code := e.INVALD_PARAMS

	if ! valid.HasErrors() {
		if models.ExistArticleByID(id){
			models.DeleteArticle(id)
			code = e.SUCCESS
		}else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _,err := range valid.Errors {
			logging.Fatal(err.Key,err.Message)
		}
	}
	util.ResponseWithJson(code,make(map[string]interface{}),c)
}

