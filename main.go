package main

import (
	"os"
	"log"
	"bufio"
	"strings"
	"regexp"
	"fmt"
)

func main() {

//	consumerKey := os.Getenv("twitter_consumer_key")
//	accessKey := os.Getenv("twitter_access_key")
//	accessToken := os.Getenv("twitter_access_token")
//	secretAccessToken := os.Getenv("twitter_access_secret_token")
//	r := gin.Default()
//
//	r.GET("/", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"test": "ok",
//		})
//	})
//
//	r.GET("/collect", func(c *gin.Context) {
//		collectTweets(consumerKey, accessKey, accessToken, secretAccessToken)
//		c.JSON(200, gin.H{
//			"": "",
//		})
//	})
//
//	r.Run()

	file, err := os.Open("/Users/tylerpachal/Golang/src/github.com/tylerpachal/GoPolitics/textfiles/example_tweets_text.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()


	bernieRegexp := makeRegexp(BernieAliases)
	donaldRegexp := makeRegexp(DonaldAliases)
	hillaryRegexp := makeRegexp(HillaryAliases)
	tedRegexp := makeRegexp(TedAliases)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        text := scanner.Text()

		// If the tweet is too short, get rid of it
		if len(text) < 60 { continue }

		// If the tweet has a link, ignore it
		if strings.Contains(text, "http") { continue }

		// If the tweet is part of a series (part 1 of 2, etc) ignore it
		if strings.HasSuffix(text, "…") { continue }

		// If the first two tokens are handles, gonna assume its just a list of handles, and we're gonna ignore it
		tokens := strings.Split(text, " ")
		if len(tokens) < 2 { continue }
		if strings.HasPrefix(tokens[0], "@") && strings.HasPrefix(tokens[1], "@") { continue }

		// If it is a re-tweet, remove the RT syntax
		if (strings.HasPrefix(text, "RT")) {
			slices := strings.SplitAfterN(text, ": ", 2)
			text = slices[len(slices)-1]
		}

		// Re-tokenize
		tokens = strings.Split(text, " ")

		answer := ""

		// Replace all occurrences of whichever candidate is mentioned first
		for _, token := range tokens {
			if bernieRegexp.MatchString(token) {
				text = bernieRegexp.ReplaceAllString(text, "______")
				answer = "bernie"
				break
			}
			if hillaryRegexp.MatchString(token) {
				text = hillaryRegexp.ReplaceAllString(text, "______")
				answer = "hillary"
				break
			}
			if tedRegexp.MatchString(token) {
				text = tedRegexp.ReplaceAllString(text, "______")
				answer = "ted"
				break
			}
			if donaldRegexp.MatchString(token) {
				text = donaldRegexp.ReplaceAllString(text, "______")
				answer = "donald"
				break
			}
		}

		// If there was answer, the tweet wasn't about anyone, and toss it
		if len(answer) < 1 {
			continue
		}

		println("Question:", text)
		println("Answer:", answer)

		println()
	}

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func makeRegexp(strings []string) *regexp.Regexp {
	s := "(?i)"
	for i, a := range strings {
		seperator := "|"
		if i == 0 {
			seperator = ""
		}
		s = fmt.Sprintf("%s%s%s", s, seperator, a)
	}
	return regexp.MustCompile(s)
}