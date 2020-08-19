package flag

import (
	"flag"

	dTwitter "github.com/tokopedia/dev/compfest/delivery/twitter"
	uTwitter "github.com/tokopedia/dev/compfest/usecase/twitter"
)

type delivery struct {
	usecase uTwitter.Usecase
}

func New(usecase uTwitter.Usecase) dTwitter.DeliveryFlag {
	handler := &delivery{
		usecase: usecase,
	}

	// define available flag command
	tweet := flag.String("tweet", "", "tweet something")

	// parse command
	flag.Parse()

	// post based on flag
	handler.PostTweet(*tweet)

	return &delivery{
		usecase: usecase,
	}
}
