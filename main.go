package main

import (
	"gitlab_webhook/wecom"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routerWecom := r.Group("/wecom")
	{
		routerWecom.POST("/ping", wecom.Ping)
	}

	r.Run(":3000")
}
