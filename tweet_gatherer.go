package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"strconv"
)

func collectTweets(consumerKey, accessKey, accessToken, accessTokenSecret string) {

	const TRUMP = "Trump"
	const CRUZ = "Cruz"
	const KASICH = "Kasich"
	const SANDERS = "Sanders"
	const CLINTON = "Clinton"
	const TWEETS_COUNT = 200

	var tweets [5 * TWEETS_COUNT]anaconda.Tweet
	i := 0

	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(accessKey)
	api := anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	tweetCount := url.Values{}
	tweetCount.Set("count", strconv.Itoa(TWEETS_COUNT))

	searchResultTrump, _ := api.GetSearch(TRUMP, tweetCount)
	for _, tweet := range searchResultTrump.Statuses {
		tweets[i] = tweet
		i++
	}

	searchResultCruz, _ := api.GetSearch(CRUZ, tweetCount)
	for _, tweet := range searchResultCruz.Statuses {
		tweets[i] = tweet
		i++
	}

	searchResultKasich, _ := api.GetSearch(KASICH, tweetCount)
	for _, tweet := range searchResultKasich.Statuses {
		tweets[i] = tweet
		i++
	}

	searchResultSanders, _ := api.GetSearch(SANDERS, tweetCount)
	for _, tweet := range searchResultSanders.Statuses {
		tweets[i] = tweet
		i++
	}

	searchResultClinton, _ := api.GetSearch(CLINTON, tweetCount)
	for _, tweet := range searchResultClinton.Statuses {
		tweets[i] = tweet
		i++
	}

	// Return tweets[] when ready

	println("OK")
}
