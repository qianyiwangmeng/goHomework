/**
  @author: qianyi  2022/5/17 21:51:00
  @note:
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"goHomework/dao"
	"goHomework/logic"
	"goHomework/model"
	"log"
	"net/http"
	"strconv"
)

// 获取排行榜前几
func ArticleTopNHandler(c *gin.Context) {
	nStr := c.Param("n")
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		log.Printf("获取排行榜topN数据转化失败%v")
		n = 5
	}
	articles, err := logic.GetArticleReadCountTopN(n)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "获取热点文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 1,
		Msg:  "获取热点文章成功",
		Data: articles,
	})
}

// 分页加载文章
func LoadArticlesHandler(c *gin.Context) {
	var req ReqPage
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(200, "参数绑定失败")
		return
	}
	articles, err := dao.LoadArticles(req.PageIndex, req.PageSize)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "分页加载文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "分页加载文章成功",
		Data: articles,
	})

}

// 写文章
func AddArticleHandler(c *gin.Context) {
	var req ReqArticle
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(200, "参数绑定失败")
		return
	}

	user, _ := c.Cookie("login_user")

	var article = model.Article{
		Title:   req.Title,
		Tags:    req.Tags,
		Short:   req.Short,
		Content: req.Content,
		Author:  user,
	}
	err = dao.AddArticles(&article)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "创建文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "创建文章成功",
	})
}

// 根据id删除文章
func DeleteArticleHandler(c *gin.Context) {
	query, ok := c.GetQuery("id")
	if !ok {
		c.String(200, "参数绑定失败")
		return
	}

	id, _ := strconv.Atoi(query)
	err := dao.DeleteArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "删除文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "删除文章成功",
	})

}

// 修改文章
func UpdateArticleHandler(c *gin.Context) {
	var req ReqUpdateArticle
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.String(200, "参数绑定失败")
		return
	}

	var updateArticle = model.Article{
		ID:      req.ID,
		Title:   req.Title,
		Tags:    req.Tags,
		Short:   req.Short,
		Content: req.Content,
	}
	err = dao.UpdateArticles(&updateArticle)
	if err != nil {
		log.Printf("更新错误：%v", err)
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "更新文章失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "更新文章成功",
	})

}

// 根据id查询某一篇文章
func ShowArticle(c *gin.Context) {
	query, ok := c.GetQuery("id")
	if !ok {
		c.String(200, "参数绑定失败")
		return
	}

	id, _ := strconv.Atoi(query)
	article, err := dao.FindArticle(id)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "获取文章失败",
		})
		return
	}

	if article.ID == 0 {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "未找到该篇文章",
		})
		return
	}

	err = logic.IncArticleReadCount(query)
	if err != nil {
		log.Printf("文章阅读加一失败：%v", err)
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "获取文章成功",
		Data: article,
	})
}
