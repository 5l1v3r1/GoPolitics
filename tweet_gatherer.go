package main

import (
	"fmt"
	"github.com/ChimeraCoder/anaconda"
)

func collectTweets(consumerKey, accessKey, accessToken, accessTokenSecret string) {

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(accessKey)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	searchResult, _ := api.GetSearch("Trump", nil)
	for _, tweet := range searchResult.Statuses {
		fmt.Println("---")
		fmt.Println(tweet.Coordinates)
		fmt.Println(tweet.Latitude())
		fmt.Println(tweet.Longitude())
		fmt.Println(tweet.Id)
		fmt.Println(tweet.Text)
		fmt.Println(tweet.CreatedAt)
		fmt.Println(tweet.RetweetCount)
	}
}
