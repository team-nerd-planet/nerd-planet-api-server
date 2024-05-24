package entity

import "github.com/lib/pq"

type SkillTag struct {
	ID      uint           `gorm:"column:id;primarykey"`
	Name    string         `gorm:"column:name;type:varchar;not null"`
	Keyword pq.StringArray `gorm:"column:keyword;type:text[];not null"`
}

type SkillTagRepo interface {
	FindAll() ([]SkillTag, error)
}
