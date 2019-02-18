package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	casbin2 "test/test-casbin/casbin"
)

// 添加权限
func AddCasbin(c *gin.Context) {
	rolename := c.PostForm("rolename")
	path := c.PostForm("path")
	method := c.PostForm("method")
	ptype := c.PostForm("ptype")
	fmt.Println( rolename)
	casbin := casbin2.CasbinModel{
		Ptype:    ptype,
		RoleName: rolename,
		Path:     path,
		Method:   method,
	}
	isok := casbin.AddCasbin(casbin)
	if isok {
		c.JSON(http.StatusOK, gin.H{
			"SUCESS": true,
			"msg":    "保存成功",
		})
		c.Abort()
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "保存失败",
		})
		c.Abort()
		return
	}
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg": "hello world!",
	})

}
