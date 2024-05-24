package item_dto

import (
	"time"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type FindAllItemReq struct {
	Company     *string                   `form:"company" binding:"omitempty,min=1"`            // 회사 이름 검색 키워드
	CompanySize *[]entity.CompanySizeType `form:"company_size" binding:"omitempty,gte=0,lte=4"` // 회사 규모 (0:스타트업, 1:중소기업, 2:중견기업, 3:대기업, 4:외국계)
	JobTags     *[]int64                  `form:"job_tags" binding:"omitempty"`                 // 관련 직무 DB ID 배열
	SkillTags   *[]int64                  `form:"skill_tags" binding:"omitempty"`               // 관련 스킬 DB ID 배열
	Page        *int                      `form:"page" binding:"required,gte=1"`                // 페이지
}

type FindAllItemRes struct {
	ItemID          uint                   `json:"item_id"`           // 글 DB ID
	ItemTitle       string                 `json:"item_title"`        // 글 제목
	ItemDescription string                 `json:"item_description"`  // 글 설명
	ItemLink        string                 `json:"item_link"`         // 글 URL
	ItemPublished   time.Time              `json:"item_published"`    // 개시 시간
	FeedName        string                 `json:"feed_name"`         // 회사 이름
	FeedTitle       string                 `json:"feed_title"`        // 회사 Feed 제목
	FeedLink        string                 `json:"feed_link"`         // 회사 URL
	CompanySize     entity.CompanySizeType `json:"company_size"`      // 회사 규모
	JobTagIDArr     []int64                `json:"job_tags_id_arr"`   // 관련 직무 DB ID 배열
	SkillTagIDArr   []int64                `json:"skill_tags_id_arr"` // 관련 스킬 DB ID 배열
}

func NewFindAllItemRes(viewItem entity.ItemView) FindAllItemRes {
	return FindAllItemRes{
		ItemID:          viewItem.ItemID,
		ItemTitle:       viewItem.ItemTitle,
		ItemDescription: viewItem.ItemDescription,
		ItemLink:        viewItem.ItemLink,
		ItemPublished:   viewItem.ItemPublished,
		FeedName:        viewItem.FeedName,
		FeedTitle:       viewItem.FeedTitle,
		FeedLink:        viewItem.FeedLink,
		CompanySize:     viewItem.CompanySize,
		JobTagIDArr:     viewItem.JobTagIDArr,
		SkillTagIDArr:   viewItem.SkillTagIDArr,
	}
}
