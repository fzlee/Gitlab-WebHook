package wecom

import (
	"gitlab_webhook/utils"

	"github.com/gin-gonic/gin"
)

func stringInMap(a string, maps map[string]func(*gin.Context)) bool {
	for item, _ := range maps {
		if a == item {
			return true
		}
	}
	return false
}

func Ping(c *gin.Context) {
	pushTypes := c.Request.Header["X-Gitlab-Event"]
	pushType := utils.GetFirstItem(pushTypes)

	if pushType == nil {
		utils.FailedResponse(c, "empty header X-Gitlab-Event")
		return
	}

	maps := map[string]func(c *gin.Context){
		"Push Hook":          handlePushHook,
		"Merge Request Hook": handlePushHook,
		"Pipeline Hook":      handlePushHook,
		"Job Hook":           handlePushHook,
		"Deployment Hook":    handlePushHook,
		"Release Hook":       handlePushHook,
	}

	if pushType == nil || !stringInMap(*pushType, maps) {
		utils.FailedResponse(c, "unknown header X-Gitlab-Event: "+*pushType)
		return
	}

	// handle webhook
	maps[*pushType](c)

	utils.SuccessResponse(c)
}
