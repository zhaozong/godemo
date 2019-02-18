package router

import (
	"github.com/gin-gonic/gin"
	"test/test-casbin/apis"
	"test/test-casbin/middleware"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.POST("/test", apis.AddCasbin)
	authrequired := router.Group("/apis")
	authrequired.Use(middleware.AuthCheckRole())

	authrequired.POST("/login", apis.Login)

	return router
}
