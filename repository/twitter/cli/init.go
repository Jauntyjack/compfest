package twitter

import (
	rTwitter "github.com/tokopedia/dev/compfest/repository/twitter"
)

type repository struct {
}

func New() rTwitter.Repository {
	return &repository{}
}
