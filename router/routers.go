package router

import (
	"github.com/gin-gonic/gin"
	"test2/middleware/jwt"
	"test2/pkg/setting"
	v1 "test2/router/v1"
)

func InitRouter() *gin.Engine {
	c := gin.New()
	c.Use(gin.Logger())
	c.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	c.GET("/auth", v1.GetAuth)

	apiv1 := c.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTags)
		apiv1.POST("/tags/:id/", v1.UpdateTags)
		apiv1.DELETE("/tags/:id/", v1.DeleteTags)

		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.ArticleDetail)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.POST("/articles/:id", v1.UpdateArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return c
}
