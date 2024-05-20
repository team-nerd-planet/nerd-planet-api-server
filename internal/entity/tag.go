package entity

type Tag struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar;not null"`
}
