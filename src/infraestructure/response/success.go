package response

import "github.com/gin-gonic/gin"

func SendSucess(c *gin.Context, msg string, status int, data interface{}) {
	c.JSON(status, gin.H{"data": data, "error": "", "status": status, "message": msg})
}
