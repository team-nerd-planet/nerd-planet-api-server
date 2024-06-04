package feed

import (
	"log/slog"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type FeedUsecase struct {
	feedRepo entity.FeedRepo
}

func NewFeedUsecase(feedRepo entity.FeedRepo) FeedUsecase {
	return FeedUsecase{
		feedRepo: feedRepo,
	}
}

func (fu FeedUsecase) FindAll(keyword *string) (*[]entity.Feed, bool) {
	feeds, err := fu.feedRepo.FindAll(keyword)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return &feeds, true
}
