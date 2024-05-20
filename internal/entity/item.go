package entity

import (
	"time"

	"github.com/lib/pq"
)

type Item struct {
	ID          uint      `gorm:"column:id;primarykey"`
	Title       string    `gorm:"column:title;type:varchar;not null"`
	Description string    `gorm:"column:description;type:varchar;not null"`
	Link        string    `gorm:"column:link;type:varchar;not null"`
	Updated     time.Time `gorm:"column:updated;type:timestamp;not null"`
	GUID        string    `gorm:"column:guid;type:varchar;not null"`
	FeedID      uint      `gorm:"column:feed_id;type:int8;not null"`
	Tags        []Tag     `gorm:"many2many:item_tags;"`
}

type ViewItem struct {
	ItemID          uint            `gorm:"column:item_id;type:int8"`
	ItemTitle       string          `gorm:"column:item_title;type:varchar"`
	ItemDescription string          `gorm:"column:item_description;type:varchar"`
	ItemLink        string          `gorm:"column:item_link;type:varchar"`
	ItemUpdated     time.Time       `gorm:"column:item_updated;type:timestamp"`
	FeedName        string          `gorm:"column:feed_name;type:varchar"`
	FeedTitle       string          `gorm:"column:feed_title;type:varchar"`
	FeedLink        string          `gorm:"column:feed_link;type:varchar"`
	CompanySize     CompanySizeType `gorm:"column:company_size;type:int2"`
	TagIDArr        pq.Int64Array   `gorm:"column:tag_id_arr;type:int8[]"`
}

func (ViewItem) TableName() string {
	return "vw_items"
}

type ItemRepo interface {
	CountView(companySizeArr *[]CompanySizeType, itemTagIDArr *[]int64) (int64, error)
	FindAllView(companySizeArr *[]CompanySizeType, itemTagIDArr *[]int64, perPage int, page int) ([]ViewItem, error)
}
