package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleResponseWithStatus(c *gin.Context, status int, data interface{}, err error) {
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}
	c.JSON(status, data)
}

func HandleResponse(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
