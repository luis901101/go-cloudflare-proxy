package main

import (
	"cloudflare-proxy/auth"
	//"cloudflare-proxy/auth"
	"cloudflare-proxy/conf"
	"cloudflare-proxy/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

var version = "0.0.1+1"
var db = make(map[string]string)

func setupEngine() *gin.Engine {
	config := conf.Load()

	// Disable Console Color
	// gin.DisableConsoleColor()
	engine := gin.Default()
	protectedRoutes := engine.Group("/")
	protectedRoutes.Use(auth.Middleware(config))

	// Initialize controllers
	imageController := controller.NewImageController(config)
	streamController := controller.NewStreamController(config)

	// Register controller routes
	imageController.RegisterRoutes(protectedRoutes)
	streamController.RegisterRoutes(protectedRoutes)

	//
	//// Ping test
	//engine.GET("/ping", func(c *gin.Context) {
	//	c.String(http.StatusOK, "pong")
	//})
	//
	//// Get user value
	//engine.GET("/user/:name", func(c *gin.Context) {
	//	user := c.Params.ByName("name")
	//	value, ok := db[user]
	//	if ok {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	//	} else {
	//		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	//	}
	//})
	//
	//// Authorized group (uses gin.BasicAuth() middleware)
	//// Same than:
	//// authorized := engine.Group("/")
	//// authorized.Use(gin.BasicAuth(gin.Credentials{
	////	  "foo":  "bar",
	////	  "manu": "123",
	////}))
	//authorized := engine.Group("/", gin.BasicAuth(gin.Accounts{
	//	"foo":  "bar", // user:foo password:bar
	//	"manu": "123", // user:manu password:123
	//}))
	//
	///* example curl for /admin with basicauth header
	//   Zm9vOmJhcg== is base64("foo:bar")
	//
	//	curl -X POST \
	//  	http://localhost:8080/admin \
	//  	-H 'authorization: Basic Zm9vOmJhcg==' \
	//  	-H 'content-type: application/json' \
	//  	-d '{"value":"bar"}'
	//*/
	//authorized.POST("admin", func(c *gin.Context) {
	//	user := c.MustGet(gin.AuthUserKey).(string)
	//
	//	// Parse JSON
	//	var json struct {
	//		Value string `json:"value" binding:"required"`
	//	}
	//
	//	if c.Bind(&json) == nil {
	//		db[user] = json.Value
	//		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	//	}
	//})

	return engine
}

func main() {
	fmt.Printf("Server running version: %s\n", version)
	engine := setupEngine()
	// Listen and Server in 0.0.0.0:8080
	err := engine.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %e", err)
		return
	}
}
