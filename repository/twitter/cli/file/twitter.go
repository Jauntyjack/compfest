package twitter

import (
	"log"
	"os"
)

func (r *repository) PostTweet(tweet string) {
	log.Print("posting on twitter file!")
	f, err := os.OpenFile("tweet.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Print(err)
	}
	defer f.Close()
	if _, err := f.WriteString(tweet + "\n"); err != nil {
		log.Print(err)
	}
}
