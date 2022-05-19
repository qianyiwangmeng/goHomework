/**
  @author: qianyi  2022/5/18 12:41:00
  @note:
*/
package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BasicAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := c.Cookie("login_user")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"msg": "登录失效，请重新登录",
			})
			c.Abort()
			return
		}
		// 有cookie则验证通过
		c.Next()

	}
}
