package entity

import "github.com/lib/pq"

type JobTag struct {
	ID      uint           `gorm:"column:id;primarykey"`
	Name    string         `gorm:"column:name;type:varchar;not null"`
	Keyword pq.StringArray `gorm:"column:keyword;type:text[];not null"`
}

type JobTagRepo interface {
	FindAll() ([]JobTag, error)
}
