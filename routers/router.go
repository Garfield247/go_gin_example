package routers

import (
	"github.com/Garfield247/go_gin_example/middleware/jwt"
	"github.com/Garfield247/go_gin_example/pkg/setting"
	"github.com/Garfield247/go_gin_example/routers/api"
	v1 "github.com/Garfield247/go_gin_example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"MSG":"test",
		})
	})
	//AUTH
	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/V1")
	apiV1.Use(jwt.JWT())
	{
		//tag
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		//article
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.EditArticle)
		apiV1.DELETE("articles/:id", v1.DeleteArticle)

	}
	return r
 }