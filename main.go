package main

import (
	"net/http"
	"seaals-api/seaals"

	"github.com/gin-gonic/gin"
	"gopkg.in/gographics/imagick.v3/imagick"
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

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "good"})
		}
	})

	// Seaal stuff

	r.GET("/seal", func(c *gin.Context) {
		sm := NewSealMagick()
		// Load the image
		sm.Mw.ReadImage(seaals.GetRandomSeal())
		// Get bytes and return
		b := sm.GetImageBytes()
		c.Data(http.StatusOK, "image/jpeg", b)
	})

	r.GET("/seal/says/:text", func(c *gin.Context) {
		sm := NewSealMagick()
		text := c.Params.ByName("text")
		// Load the image
		sm.Mw.ReadImage(seaals.GetRandomSeal())
		// Draw text
		sm.ThickOutline(text, "Adwaita-Mono", imagick.GRAVITY_SOUTH)
		sm.DrawCurrent()
		// Get bytes and return
		b := sm.GetImageBytes()
		c.Data(http.StatusOK, "image/jpeg", b)
	})

	return r
}

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
