package module

import (
	rTwitter "github.com/tokopedia/dev/compfest/repository/twitter"
	uTwitter "github.com/tokopedia/dev/compfest/usecase/twitter"
)

type twitterUseCase struct {
	repo rTwitter.Repository
}

func New(twitterFileRepo rTwitter.Repository) uTwitter.Usecase {
	return &twitterUseCase{
		repo: twitterFileRepo,
	}
}
