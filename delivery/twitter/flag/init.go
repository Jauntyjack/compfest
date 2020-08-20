package flag

import (
	"flag"

	dTwitter "github.com/jauntyjack/compfest/delivery/twitter"
	uTwitter "github.com/jauntyjack/compfest/usecase/twitter"
)

type delivery struct {
	usecase uTwitter.Usecase
}

func New(usecase uTwitter.Usecase) dTwitter.DeliveryFlag {
	handler := &delivery{
		usecase: usecase,
	}

	// define available flag command
	tweet := flag.String("tweet", "", "Allows user to tweet using via cli")

	// parse command
	flag.Parse()

	// post based on flag
	handler.PostTweet(*tweet)

	return &delivery{
		usecase: usecase,
	}
}
