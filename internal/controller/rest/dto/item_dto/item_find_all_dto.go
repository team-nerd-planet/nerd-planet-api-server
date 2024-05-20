package item_dto

import (
	"time"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type FindAllItemReq struct {
	CompanySize *[]entity.CompanySizeType `form:"company_size" binding:"required"`
	Tags        *[]int64                  `form:"tags" binding:"required"`
	Page        *int                      `form:"page" binding:"required,gte=1"`
}

type FindAllItemRes struct {
	ItemID          uint                   `json:"item_id"`
	ItemTitle       string                 `json:"item_title"`
	ItemDescription string                 `json:"item_description"`
	ItemLink        string                 `json:"item_link"`
	ItemUpdated     time.Time              `json:"item_updated"`
	FeedName        string                 `json:"feed_name"`
	FeedTitle       string                 `json:"feed_title"`
	FeedLink        string                 `json:"feed_link"`
	CompanySize     entity.CompanySizeType `json:"company_size"`
	ItemTags        []int64                `json:"item_tags"`
}

func NewFindAllItemRes(viewItem entity.ViewItem) FindAllItemRes {
	return FindAllItemRes{
		ItemID:          viewItem.ItemID,
		ItemTitle:       viewItem.ItemTitle,
		ItemDescription: viewItem.ItemDescription,
		ItemLink:        viewItem.ItemLink,
		ItemUpdated:     viewItem.ItemUpdated,
		FeedName:        viewItem.FeedName,
		FeedTitle:       viewItem.FeedTitle,
		FeedLink:        viewItem.FeedLink,
		CompanySize:     viewItem.CompanySize,
		ItemTags:        viewItem.TagIDArr,
	}
}
