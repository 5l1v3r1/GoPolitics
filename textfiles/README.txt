To run (terminal):
    go run *.go

////

Set ENV Variables in your bashrc or bash_profile

export twitter_consumer_key="key1"
export twitter_access_key="ke2"
export twitter_access_token="ke3"
export twitter_access_secret_token="key4"

////

To store tweets to file in tweet_gatherer, add:

	// Store tweets to file
	fileHandle, _ := os.Create("textfiles/example_tweets_text.txt")
	writer := bufio.NewWriter(fileHandle)
	defer fileHandle.Close()

	for j := 0; j < len(tweets); j++ {
		fmt.Fprintln(writer, tweets[j].Text)
		fmt.Fprintln(writer, "")
		writer.Flush()
	}