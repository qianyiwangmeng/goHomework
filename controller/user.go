/**
  @author: qianyi  2022/5/17 21:51:00
  @note:
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"goHomework/dao"
	"goHomework/model"
	"goHomework/utils"
	"net/http"
)

// 用户注册
func RegisterHandler(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")

	usr, _ := dao.FindUser(name)

	if usr.ID != 0 {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "用户名已存在",
		})
		return
	}

	md5Password := utils.MD5(password)
	user := model.User{Name: name, Password: md5Password}

	err := dao.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "注册失败",
		})
		return
	}

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "注册成功",
	})

}

// 用户登录
func LoginHandler(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	user, _ := dao.FindUser(name)
	if user.ID == 0 {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "用户名不存在",
		})
		return
	}

	md5Password := utils.MD5(password)
	if user.Password != md5Password {
		c.JSON(http.StatusOK, &Res{
			Code: 1,
			Msg:  "密码错误",
		})
		return
	}

	// 设置cookie
	c.SetCookie("login_user", user.Name, 3600, "/", "localhost", false, true)

	c.JSON(http.StatusOK, &Res{
		Code: 0,
		Msg:  "登录成功",
	})

}
