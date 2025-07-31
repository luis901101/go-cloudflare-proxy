package auth

import (
	"cloudflare-proxy/conf"
	"cloudflare-proxy/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Middleware creates gin middleware for authorization.
func Middleware(config conf.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the Authorization header.
		authHeader := c.GetHeader(utils.HeaderApiKey)
		if authHeader == "" {
			// Stop processing and send an error response.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("%s header is required", utils.HeaderApiKey)})
			return
		}
		if authHeader != config.ApiKey {
			// Stop processing and send an error response.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid %s header", utils.HeaderApiKey)})
			return
		}

		//// 2. Validate the token format (e.g., "Bearer <token>").
		//parts := strings.Split(authHeader, " ")
		//if len(parts) != 2 || parts[0] != "Bearer" {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
		//	return
		//}
		//
		//// 3. Check if the token is valid.
		//if parts[1] != config.ApiKey {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
		//	return
		//}

		log.Println("Authorization successful")
		// If the token is valid, proceed to the next handler in the chain.
		c.Next()
	}
}
