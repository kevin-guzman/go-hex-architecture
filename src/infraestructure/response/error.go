package response

import "github.com/gin-gonic/gin"

func SendError(c *gin.Context, err, msg string, status int) {
	c.JSON(status, gin.H{"data": nil, "error": err, "status": status, "message": msg})
}
