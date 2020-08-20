package twitter

import (
	rTwitter "github.com/jauntyjack/compfest/repository/twitter"
)

type repository struct {
}

func New() rTwitter.Repository {
	return &repository{}
}
