package module

import (
	rTwitter "github.com/jauntyjack/compfest/repository/twitter"
	uTwitter "github.com/jauntyjack/compfest/usecase/twitter"
)

type twitterUseCase struct {
	repo rTwitter.Repository
}

func New(twitterFileRepo rTwitter.Repository) uTwitter.Usecase {
	return &twitterUseCase{
		repo: twitterFileRepo,
	}
}
