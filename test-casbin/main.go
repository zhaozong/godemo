package main

import (
	"test/test-casbin/router"
)

func main() {
	route := router.InitRouter()
	route.Run(":9000")
	//e := casbin.Casbin()
	//pomission := e.GetImplicitPermissionsForUser("jack")
	//fmt.Printf("pomission is %+v\n", pomission)
}
