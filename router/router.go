/**
  @author: qianyi  2022/5/17 21:44:00
  @note:
*/
package router

import (
	"github.com/gin-gonic/gin"
	"goHomework/controller"
	"goHomework/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 无需认证
	r.POST("/register", controller.RegisterHandler)
	r.POST("/login", controller.LoginHandler)
	r.GET("/article/top/:n", controller.ArticleTopNHandler)

	basicAuthGroup := r.Group("/", middleware.BasicAuth())

	article := basicAuthGroup.Group("/article")
	{
		article.POST("/load", controller.LoadArticlesHandler)
		article.POST("/add", controller.AddArticleHandler)
		article.GET("/delete", controller.DeleteArticleHandler)
		article.POST("/update", controller.UpdateArticleHandler)
		article.GET("/show", controller.ShowArticle)
	}

	// 相册
	basicAuthGroup.POST("/upload", controller.UploadHandler)
	basicAuthGroup.GET("/album", controller.ShowAlbumHandler)

	return r
}
