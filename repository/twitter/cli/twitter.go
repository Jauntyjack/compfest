package twitter

import "log"

func (r *repository) PostTweet(tweet string) {
	log.Print("Posting on twitter! : ", tweet)
}
