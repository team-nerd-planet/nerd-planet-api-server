package entity

import "time"

type CompanySizeType int

const (
	STARTUP CompanySizeType = iota //스타트업
	SMALL                          //중소기업
	MEDIUM                         //중견기업
	LARGE                          //대기업
	FOREIGN                        //외국계
)

type Feed struct {
	ID          uint            `gorm:"column:id;primarykey"`
	Name        string          `gorm:"column:name;type:varchar;not null"`
	Title       string          `gorm:"column:title;type:varchar;not null"`
	Description string          `gorm:"column:description;type:varchar;not null"`
	Link        string          `gorm:"column:link;type:varchar;not null;unique"`
	Updated     time.Time       `gorm:"column:updated;type:timestamp;not null"`
	Copyright   string          `gorm:"column:copyright;type:varchar;not null"`
	CompanySize CompanySizeType `gorm:"column:company_size;type:int2;not null"`
	RssID       uint            `gorm:"column:rss_id;type:int8;not null"`
	Items       []Item          `gorm:"foreignKey:FeedID"`
}
