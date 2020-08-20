package http

import (
	"net/http"

	dTwitter "github.com/jauntyjack/compfest/delivery/twitter"
	uTwitter "github.com/jauntyjack/compfest/usecase/twitter"
)

type delivery struct {
	usecase uTwitter.Usecase
}

func New(usecase uTwitter.Usecase) dTwitter.Delivery {
	handler := &delivery{
		usecase: usecase,
	}

	http.HandleFunc("/post", handler.PostTweet)
	return &delivery{
		usecase: usecase,
	}
}
