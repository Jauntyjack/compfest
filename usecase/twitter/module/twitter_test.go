package module

import (
	"testing"

	"github.com/golang/mock/gomock"
	cliMock "github.com/jauntyjack/compfest/repository/twitter/mocks"
)

func Test_twitterUseCase_PostTweet(t *testing.T) {

	controller := gomock.NewController(t)
	defer controller.Finish()

	repoMock := cliMock.NewMockRepository(controller)

	repoMock.EXPECT().PostTweet(gomock.Any())

	u := New(repoMock)

	t.Run("input is valid", func(t *testing.T) {
		result := u.PostTweet("halo")
		if result != "halo" {
			t.Fail()
		}
	})

	t.Run("character is more than 30", func(t *testing.T) {
		result := u.PostTweet("halohalohalohalohalohalohalohalohalohalohalohalohalohalohalohalo")
		if result != "You can't tweet more than 30 character" {
			t.Fail()
		}
	})
}
