package handler

import (
	"cloudflare-proxy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleResponseWithStatus(c *gin.Context, status int, data interface{}, err error) {
	if err != nil {
		c.AbortWithStatusJSON(status, utils.NullCheck(data, gin.H{"error": err.Error()}))
		return
	}
	c.JSON(status, data)
}

func HandleResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.NullCheck(data, gin.H{"error": err.Error()}))
		return
	}
	c.JSON(http.StatusOK, data)
}
