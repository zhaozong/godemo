package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/test-casbin/casbin"
)

func AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/test" {
			c.Next()
		} else {
			e := casbin.Casbin()
			fmt.Println(c.Request.URL.Path)
			fmt.Println(c.Request.Method)
			res, err := e.EnforceSafe("admin", c.Request.URL.Path, c.Request.Method)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status": -1,
					"msg":    "错误 " + err.Error(),
				})
				c.Abort()
				return
			}
			if res {
				c.Next()
			} else {
				c.JSON(http.StatusOK, gin.H{
					"status": 0,
					"msg":    "没有权限",
				})
				c.Abort()
				return
			}
		}
	}
}
