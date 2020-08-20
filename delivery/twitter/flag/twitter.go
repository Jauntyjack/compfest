package flag

import "log"

func (d *delivery) PostTweet(tweet string) {

	usecaseResponse := d.usecase.PostTweet(tweet)
	log.Print("Delivery output: ", usecaseResponse)
}
