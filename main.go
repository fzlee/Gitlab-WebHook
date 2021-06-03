package main

import (
	"gitlab_webhook/utils"
	"gitlab_webhook/wecom"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routerWecom := r.Group("/wecom")
	{
		routerWecom.POST("/ping", wecom.Ping)
	}

	Token := os.Getenv("WECOM_TOKEN")
	if Token == "" {
		log.Fatal("WECOM_TOKEN 为空")
		return
	}

	utils.WECOM_TOKEN = &Token

	r.Run(":3000")
}
