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
			fmt.Println(c.PostForm("role"))
			fmt.Println(c.PostForm("obj"))
			fmt.Println(c.PostForm("act"))
			res := e.Enforce(c.PostForm("role"), c.PostForm("obj"), c.PostForm("act"))
			//if err != nil {
			//	c.JSON(http.StatusInternalServerError, gin.H{
			//		"status": -1,
			//		"msg":    "错误 " + err.Error(),
			//	})
			//	c.Abort()
			//	return
			//}
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
