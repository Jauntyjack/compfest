package http

import (
	"net/http"

	dTwitter "github.com/tokopedia/dev/compfest/delivery/twitter"
	uTwitter "github.com/tokopedia/dev/compfest/usecase/twitter"
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
