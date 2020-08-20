package usecase

type Usecase interface {
	PostTweet(tweet string) string
}
