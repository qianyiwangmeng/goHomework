/**
  @author: qianyi  2022/5/18 18:49:00
  @note:
*/
package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goHomework/dao"
	"goHomework/model"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func UploadHandler(c *gin.Context) {
	file, err := c.FormFile("upload")
	if err != nil {
		c.String(200, "绑定参数失败")
		return
	}

	now := time.Now()
	fileExt := path.Ext(file.Filename)

	var fileType string
	if fileExt == ".jpg" || fileExt == ".png" || fileExt == ".gif" || fileExt == ".jpeg" {
		fileType = "img"
	}

	// 文件创建的路径
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())

	// 创建文件夹
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		log.Printf("创建文件夹出错%v", err)
		c.String(200, "服务器繁忙，请稍后重试")
		return
	}

	// 文件路径和文件名
	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
	filePathStr := path.Join(fileDir, filename)

	// 将浏览器客户端上传的文件拷贝到本地路径文件里面，此处也可以使用io操作
	err = c.SaveUploadedFile(file, filePathStr)
	if err != nil {
		log.Printf("文件路径保存失败：%v", err)
	}

	user, _ := c.Cookie("login_user")

	if fileType == "img" {
		album := &model.Album{Filepath: filePathStr, Filename: filename, CreateUser: user}
		if err = dao.AddAlbum(album); err != nil {
			c.JSON(http.StatusOK, &Res{
				Code: 1,
				Msg:  "上传图片失败",
			})
			return
		}
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "上传图片成功",
	})

}

// 展示所有上传的图片
func ShowAlbumHandler(c *gin.Context) {
	loginUser, _ := c.Cookie("login_user")
	albums, err := dao.ShowAlbum(loginUser)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "获取图片失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "获取图片成功",
		Data: albums,
	})
}
