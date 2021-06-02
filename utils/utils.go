package utils

import "github.com/gin-gonic/gin"

func GetFirstItem(items []string) *string {
	if len(items) == 0 {
		return nil
	}
	return &items[0]
}

func FailedResponse(c *gin.Context, message string) {
	c.JSON(200, gin.H{
		"sucecss": false,
		"message": message,
	})
}

func SuccessResponse(c *gin.Context) {

	c.JSON(200, gin.H{
		"sucecss": true,
	})
}
