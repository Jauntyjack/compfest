package twitter

import "net/http"

type Delivery interface {
	PostTweet(w http.ResponseWriter, req *http.Request)
}

type DeliveryFlag interface {
	PostTweet(tweet string)
}
