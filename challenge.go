package main

import "github.com/ChimeraCoder/anaconda"

type Challenge struct {
	id       string   `json: id` // This is the tweet id
	tweet    Tweet    `json: tweet`
	user     User     `json: user`
	question string   `json: question`
	answer   string   `json: answer`
	options  []string `json: options`
}

func NewChallenge(tweet anaconda.Tweet) *Challenge {
	c := Challenge{}
	c.id = tweet.IdStr

	// Create inner tweet which is subset of original tweet
	c.tweet = Tweet{
		tweet.CreatedAt,
		tweet.Text,
	}

	// Create the inner user which is a subset of the original user who tweeted
	c.user = User{
		tweet.User.ScreenName,
		tweet.User.ProfileImageURL,
		tweet.User.URL,
		tweet.User.Name,
	}

	// Create the question

	return &c
}

