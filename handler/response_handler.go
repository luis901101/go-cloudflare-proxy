package handler

import (
	"github.com/gin-gonic/gin"
)

func HandleResponseWithStatus(c *gin.Context, status int, data interface{}, err error) {
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{
			"error": err.Error(),
			"data":  data,
		})
		return
	}
	c.JSON(status, data)
}
