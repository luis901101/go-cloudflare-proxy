package main

import (
	"cloudflare-proxy/dto"
	"cloudflare-proxy/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Ping test
	r.GET("/image", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.ImageResponseDTO{
			Success: utils.BoolPtr(true),
			Result: &dto.ImageDTO{
				ID:       utils.StringPtr("12345"),
				Filename: utils.StringPtr("image.png"),
				Variants: &[]string{"thumbnail", "best"},
				Metadata: &map[string]interface{}{"asd": "value"},
				Uploaded: utils.TimeToUTCTimePtr(time.Now()),
			},
		})
	})

	r.GET("/stream", func(c *gin.Context) {
		c.JSON(http.StatusOK, dto.StreamResponseDTO{
			Success: utils.BoolPtr(true),
			Result: &dto.StreamDTO{
				Uid:     utils.StringPtr("asd-qwe-zxc"),
				Created: utils.TimeToUTCTimePtr(time.Now()),
			},
		})
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/* example curl for /admin with basicauth header
	   Zm9vOmJhcg== is base64("foo:bar")

		curl -X POST \
	  	http://localhost:8080/admin \
	  	-H 'authorization: Basic Zm9vOmJhcg==' \
	  	-H 'content-type: application/json' \
	  	-d '{"value":"bar"}'
	*/
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
