package item_dto

import (
	"time"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type FindAllItemReq struct {
	Company     *string                   `url:"company" validate:"omitempty,min=1"`                 // 회사 이름 검색 키워드
	CompanySize *[]entity.CompanySizeType `url:"company_size" validate:"omitempty,dive,gte=0,lte=4"` // 회사 규모 (0:스타트업, 1:중소기업, 2:중견기업, 3:대기업, 4:외국계)
	JobTags     *[]int64                  `url:"job_tags" validate:"omitempty,dive,gte=1"`           // 관련 직무 DB ID 배열
	SkillTags   *[]int64                  `url:"skill_tags" validate:"omitempty,dive,gte=1"`         // 관련 스킬 DB ID 배열
	Page        *int                      `url:"page" validate:"required,gte=1"`                     // 페이지
}

type FindAllItemRes struct {
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

func NewFindAllItemRes(itemView entity.ItemView) FindAllItemRes {
	return FindAllItemRes{
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
