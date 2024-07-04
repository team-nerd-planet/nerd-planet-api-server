package item_dto

import (
	"time"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type FindNextReq struct {
	ExcludedIds []int64 `url:"excluded_ids"`              // 제외할 글 ID 목록
	Limit       int32   `url:"limit" validate:"required"` // 글 목록 개수 제한
}

type FindNextRes struct {
	ItemID          uint                   `json:"item_id"`           // 글 DB ID
	ItemTitle       string                 `json:"item_title"`        // 글 제목
	ItemDescription string                 `json:"item_description"`  // 글 설명
	ItemLink        string                 `json:"item_link"`         // 글 URL
	ItemThumbnail   *string                `json:"item_thumbnail"`    // 글 썸네일
	ItemPublished   time.Time              `json:"item_published"`    // 글 개시 시간
	ItemSummary     *string                `json:"item_summary"`      // 글 요약 내용
	ItemViews       uint                   `json:"item_views"`        // 조회 수
	ItemLikes       uint                   `json:"item_likes"`        // 좋아요 수
	FeedName        string                 `json:"feed_name"`         // 회사 이름
	FeedTitle       string                 `json:"feed_title"`        // 회사 Feed 제목
	FeedLink        string                 `json:"feed_link"`         // 회사 URL
	CompanySize     entity.CompanySizeType `json:"company_size"`      // 회사 규모
	JobTagIDArr     []int64                `json:"job_tags_id_arr"`   // 관련 직무 DB ID 배열
	SkillTagIDArr   []int64                `json:"skill_tags_id_arr"` // 관련 스킬 DB ID 배열
}

func NewFindNextRes(itemViews []entity.ItemView) []FindNextRes {
	res := make([]FindNextRes, len(itemViews))

	for i, itemView := range itemViews {
		res[i] = FindNextRes{
			ItemID:          itemView.ItemID,
			ItemTitle:       itemView.ItemTitle,
			ItemDescription: itemView.ItemDescription,
			ItemLink:        itemView.ItemLink,
			ItemThumbnail:   itemView.ItemThumbnail,
			ItemPublished:   itemView.ItemPublished,
			ItemSummary:     itemView.ItemSummary,
			ItemViews:       itemView.ItemViews,
			ItemLikes:       itemView.ItemLikes,
			FeedName:        itemView.FeedName,
			FeedTitle:       itemView.FeedTitle,
			FeedLink:        itemView.FeedLink,
			CompanySize:     itemView.CompanySize,
			JobTagIDArr:     itemView.JobTagIDArr,
			SkillTagIDArr:   itemView.SkillTagIDArr,
		}
	}

	return res
}
