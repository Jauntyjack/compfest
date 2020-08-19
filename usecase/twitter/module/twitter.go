package module

func (u *twitterUseCase) PostTweet(tweet string) string {

	// check if tweet size has reach max character
	if len(tweet) > 30 {
		return "You can't tweet more than 30 character"
	}

	// post to twitter
	u.repo.PostTweet(tweet)
	return tweet
}
