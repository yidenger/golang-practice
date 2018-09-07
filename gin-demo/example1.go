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

	r.Run(":3003")
}