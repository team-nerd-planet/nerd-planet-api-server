package entity

import (
	"time"

	"github.com/lib/pq"
)

type Item struct {
	ID          uint       `gorm:"column:id;primarykey"`
	Title       string     `gorm:"column:title;type:varchar;not null"`
	Description string     `gorm:"column:description;type:varchar;not null"`
	Link        string     `gorm:"column:link;type:varchar;not null"`
	Published   time.Time  `gorm:"column:published;type:timestamp;not null"`
	GUID        string     `gorm:"column:guid;type:varchar;not null"`
	FeedID      uint       `gorm:"column:feed_id;type:int8;not null"`
	JobTags     []JobTag   `gorm:"many2many:item_job_tags;"`
	SkillTags   []SkillTag `gorm:"many2many:item_skill_tags;"`
}

type ItemView struct {
	ItemID          uint            `gorm:"column:item_id;type:int8"`
	ItemTitle       string          `gorm:"column:item_title;type:varchar"`
	ItemDescription string          `gorm:"column:item_description;type:varchar"`
	ItemLink        string          `gorm:"column:item_link;type:varchar"`
	ItemPublished   time.Time       `gorm:"column:item_published;type:timestamp"`
	FeedName        string          `gorm:"column:feed_name;type:varchar"`
	FeedTitle       string          `gorm:"column:feed_title;type:varchar"`
	FeedLink        string          `gorm:"column:feed_link;type:varchar"`
	CompanySize     CompanySizeType `gorm:"column:company_size;type:int2"`
	JobTagIDArr     pq.Int64Array   `gorm:"column:job_tags_id_arr;type:int8[]"`
	SkillTagIDArr   pq.Int64Array   `gorm:"column:skill_tags_id_arr;type:int8[]"`
}

func (ItemView) TableName() string {
	return "vw_items"
}

type ItemRepo interface {
	CountView(company *string, companySizes *[]CompanySizeType, jobTags, skillTags *[]int64) (int64, error)
	FindAllView(company *string, companySizes *[]CompanySizeType, jobTags, skillTags *[]int64, perPage int, page int) ([]ItemView, error)
}
