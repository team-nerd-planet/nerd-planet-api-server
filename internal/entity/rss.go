package entity

import "time"

type Rss struct {
	ID      uint      `gorm:"column:id;primarykey"`
	Name    string    `gorm:"column:name;type:varchar;not null;unique"`
	Link    string    `gorm:"column:link;type:varchar;not null;unique"`
	Updated time.Time `gorm:"column:updated;type:timestamp"`
	Ok      bool      `gorm:"column:ok;type:bool"`
	Error   string    `gorm:"column:error;type:varchar"`
	Feed    Feed      `gorm:"foreignKey:RssID"`
}
