package main

import (
	"log"
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		fmt.Println("Hit root")
		c.HTML(http.StatusOK, "test.html", nil)
	})

	router.Run(":" + port)
}
