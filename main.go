package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {

	consumerKey := os.Getenv("twitter_consumer_key")
	accessKey := os.Getenv("twitter_access_key")
	accessToken := os.Getenv("twitter_access_token")
	secretAccessToken := os.Getenv("twitter_access_secret_token")
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"test": "ok",
		})
	})

	r.GET("/collect", func(c *gin.Context) {
		collectTweets(consumerKey, accessKey, accessToken, secretAccessToken)
		c.JSON(200, gin.H{
			"": "",
		})
	})

	tweets := collectTweets(consumerKey, accessKey, accessToken, secretAccessToken)

	println(len(tweets))
	for j := 0; j < len(tweets); j++ {
		println(tweets[j].Text)
	}

	r.Run()
}
