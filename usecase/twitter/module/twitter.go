package module

func (u *twitterUseCase) PostTweet(tweet string) string {

	// post to twitter
	u.repo.PostTweet(tweet)
	return tweet
}
