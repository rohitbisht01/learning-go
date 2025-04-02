package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rohitbisht01/url-shortener/api/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	router := gin.Default()
	setUpRouters(router)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Fatal(router.Run(":" + port))
}

func setUpRouters(router *gin.Engine) {
	router.POST("/api/v1", routes.ShortenUrl)
	router.GET("/api/v1/:shortId", routes.GetByShortId)
	router.PUT("/api/v1/:shortId", routes.EditUrl)
	router.DELETE("/api/v1/:shortId", routes.DeleteUrl)
	router.POST("/api/v1/addTag", routes.AddTag)
}
