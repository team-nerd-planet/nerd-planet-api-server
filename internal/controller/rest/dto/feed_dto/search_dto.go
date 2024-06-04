package feed_dto

import "github.com/team-nerd-planet/api-server/internal/entity"

type SearchReq struct {
	NameKeyword string `form:"name_keyword" binding:"min=1"` // 회사 이름 검색어
}

type SearchRes struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewSearchRes(feeds []entity.Feed) []SearchRes {
	res := make([]SearchRes, len(feeds))

	for i, feed := range feeds {
		res[i] = SearchRes{
			ID:   feed.ID,
			Name: feed.Name,
		}
	}

	return res
}
