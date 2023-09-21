//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, tweetChan chan *Tweet) {
	defer close(tweetChan)

	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
			// return tweets
		}

		// tweets = append(tweets, tweet)
		tweetChan <- tweet
	}
	return
}

func consumer(tweets chan *Tweet) {
	for t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	tweets := make(chan *Tweet)
	// Producer - delivers Tweet to channel
	go producer(stream, tweets)

	// Consumer - needs to receive the tweets
	consumer(tweets)

	fmt.Printf("Process took %s\n", time.Since(start))
}
