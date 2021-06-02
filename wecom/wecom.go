package wecom

import "github.com/gin-gonic/gin"

acceptedHooks = []string{
	"Push Hook",
	"Merge Request Hook",
	"Pipeline Hook",
	"Job Hook",
	"Deployment Hook",
	"Release Hook"
}

func stringInArray(a String, b []String) boolean {
	for _, item := range b {
		if a == b {
			return true
		}
	}
	return false
}


func Ping(c *gin.Context) {
	pushType := c.Request.Header["X-Gitlab-Event"]
	if stringInArray(pushType, acceptedHooks) {
		return c.JSON(200, gin.H{
			"sucecss": false,
			"message": "unknown header X-Gitlab-Event"
		})
	}
}
