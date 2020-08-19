package http

import (
	"fmt"
	"net/http"
)

func (d *delivery) PostTweet(w http.ResponseWriter, r *http.Request) {

	// read tweet parameter
	tweet := r.FormValue("tweet")

	// post tweet
	usecaseResponse := d.usecase.PostTweet(tweet)

	// write response to http writer
	response := usecaseResponse
	fmt.Fprintf(w, response)
}
