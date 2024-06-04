package rest

import (
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/feed_dto"
	"github.com/team-nerd-planet/api-server/internal/usecase/feed"
)

type FeedController struct {
	feedUcase feed.FeedUsecase
}

func NewFeedController(feedUsecase feed.FeedUsecase) FeedController {
	return FeedController{
		feedUcase: feedUsecase,
	}
}

func (fc FeedController) Search(req feed_dto.SearchReq) ([]feed_dto.SearchRes, bool) {
	res := make([]feed_dto.SearchRes, 0)

	feeds, ok := fc.feedUcase.FindAll(&req.NameKeyword)
	if !ok {
		return res, false
	}

	return feed_dto.NewSearchRes(*feeds), true
}
