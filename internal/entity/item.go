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
	Thumbnail   *string    `gorm:"column:thumbnail;type:varchar"`
	Published   time.Time  `gorm:"column:published;type:timestamp;not null"`
	GUID        string     `gorm:"column:guid;type:varchar;not null"`
	Summary     string     `gorm:"column:summary;type:varchar"`
	Views       uint       `gorm:"column:views;type:int8;not null;default:0"`
	Likes       uint       `gorm:"column:likes;type:int8;not null;default:0"`
	FeedID      uint       `gorm:"column:feed_id;type:int8;not null"`
	JobTags     []JobTag   `gorm:"many2many:item_job_tags;"`
	SkillTags   []SkillTag `gorm:"many2many:item_skill_tags;"`
}

type ItemView struct {
	ItemID          uint            `gorm:"column:item_id;type:int8"`
	ItemTitle       string          `gorm:"column:item_title;type:varchar"`
	ItemDescription string          `gorm:"column:item_description;type:varchar"`
	ItemLink        string          `gorm:"column:item_link;type:varchar"`
	ItemThumbnail   *string         `gorm:"column:item_thumbnail;type:varchar"`
	ItemPublished   time.Time       `gorm:"column:item_published;type:timestamp"`
	ItemSummary     *string         `gorm:"column:item_summary;type:varchar"`
	ItemViews       uint            `gorm:"column:item_views;type:int8"`
	ItemLikes       uint            `gorm:"column:item_likes;type:int8"`
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
	FindAllView(company *string, companySizes *[]CompanySizeType, jobTags, skillTags *[]int64, perPage, page int) ([]ItemView, error)
	Exist(id int64) (bool, error)
	IncreaseViewCount(id int64) (int64, error)
	IncreaseLikeCount(id int64) (int64, error)
	FindAllViewByExcludedIds(ids []int64, perPage int32) ([]ItemView, error)
}
