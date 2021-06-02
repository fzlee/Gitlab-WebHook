package wecom

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func handlePushHook(c *gin.Context) {
	var message ProjectMessage
	c.ShouldBindJSON(&message)

	response := NewTextResponse("测试消息")
	r, _ := json.Marshal(response)
	sendToWeCom(r)
	println("sssss")
}
