package main

import (
	"io"
	"github.com/gin-gonic/gin"
	"time"
	"os"
	"net/http"
	"fmt"
	// "encoding/json"
)

type Login struct {
	User string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Set("example", "123")

		c.Next()

		latency := time.Since(t)

		fmt.Println("time: ", latency)
	}
}

func main() {
	r := gin.New()

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.GET("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": time.Now(),
		})
	})

	
	r.POST("/login", func(c *gin.Context){
		fmt.Println(c);
		var json Login
		err := c.ShouldBindJSON(&json); 
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	r.GET("/apple", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://apple.cn")
	})

	r.Use(Logger())

	r.GET("/middleware", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		fmt.Println(example)

		c.JSON(http.StatusOK, gin.H{"message": "good job"})
	})

	r.Run(":3003")
}