package wecom

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func handlePushHook(c *gin.Context) {
	var message PushMessage
	c.ShouldBindJSON(&message)

	template := `
### gitlab通知: %s
用户: %s
项目: %s
分支: %s
`
	content := fmt.Sprintf(template, message.ObjectKind, message.UserName, message.Project.Name, message.Ref)

	response := NewMarkdownResponse(content)
	r, _ := json.Marshal(response)
	sendToWeCom(r)
	println("sssss")
}

func handlePipelineHook(c *gin.Context) {
	var message PipelineMessage
	c.ShouldBindJSON(&message)

	template := `### gitlab通知: %s
用户: %s
项目: %s
分支: %s
结果: %s
`
	content := fmt.Sprintf(template, message.ObjectKind, message.User.Name, message.Project.Name, message.ObjectAttributes.Ref, message.ObjectAttributes.Status)

	response := NewMarkdownResponse(content)
	r, _ := json.Marshal(response)
	sendToWeCom(r)
}
