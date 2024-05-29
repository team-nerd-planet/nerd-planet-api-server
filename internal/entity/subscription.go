package entity

import "github.com/lib/pq"

type Subscription struct {
	ID                      uint          `gorm:"column:id;primarykey"`
	Email                   string        `gorm:"column:email;type:varchar;not null;unique"`
	Name                    *string       `gorm:"column:name;type:varchar"`
	Division                *string       `gorm:"column:division;type:varchar"`
	PreferredCompanyArr     pq.Int64Array `gorm:"column:preferred_company_arr;type:int8[];not null"`
	PreferredCompanySizeArr pq.Int64Array `gorm:"column:preferred_company_size_arr;type:int8[];not null"`
	PreferredJobArr         pq.Int64Array `gorm:"column:preferred_job_arr;type:int8[];not null"`
	PreferredSkillArr       pq.Int64Array `gorm:"column:preferred_skill_arr;type:int8[];not null"`
}

type SubscriptionRepo interface {
	ExistEmail(email string) (*int64, error)
	Create(subscription Subscription) (*Subscription, error)
	Update(id int64, subscription Subscription) (*Subscription, error)
	Delete(id int64) (*Subscription, error)
}
