package feed_dto

import "github.com/team-nerd-planet/api-server/internal/entity"

type SearchReq struct {
	NameKeyword string `url:"name_keyword" validate:"min=1"` // 회사 이름 검색어
}

type SearchRes struct {
	ID   uint   `json:"id"`   // 검색된 회사 ID
	Name string `json:"name"` // 검색된 회사 이름
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
