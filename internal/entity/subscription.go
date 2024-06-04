package entity

import (
	"time"

	"github.com/lib/pq"
)

type Subscription struct {
	ID                      uint          `gorm:"column:id;primarykey" json:"id"`
	Email                   string        `gorm:"column:email;type:varchar;not null;unique" json:"email"`
	Name                    *string       `gorm:"column:name;type:varchar" json:"name"`
	Division                *string       `gorm:"column:division;type:varchar" json:"division"`
	PreferredCompanyArr     pq.Int64Array `gorm:"column:preferred_company_arr;type:int8[];not null" json:"preferred_company_arr"`
	PreferredCompanySizeArr pq.Int64Array `gorm:"column:preferred_company_size_arr;type:int8[];not null" json:"preferred_company_size_arr"`
	PreferredJobArr         pq.Int64Array `gorm:"column:preferred_job_arr;type:int8[];not null" json:"preferred_job_arr"`
	PreferredSkillArr       pq.Int64Array `gorm:"column:preferred_skill_arr;type:int8[];not null" json:"preferred_skill_arr"`
	Published               time.Time     `gorm:"column:published;type:timestamp;not null" json:"published"`
}

type SubscriptionRepo interface {
	ExistEmail(email string) (*int64, error)
	Create(subscription Subscription) (*Subscription, error)
	Update(id int64, subscription Subscription) (*Subscription, error)
	Delete(id int64) (*Subscription, error)
}
