package main

import "test/test-casbin/router"

func main() {
	route := router.InitRouter()
	route.Run(":9000")
}
