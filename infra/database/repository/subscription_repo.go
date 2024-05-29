package repository

import (
	"errors"

	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
	"gorm.io/gorm"
)

type SubscriptionRepo struct {
	db *database.Database
}

func NewSubscriptionRepo(db *database.Database) entity.SubscriptionRepo {
	db.AutoMigrate(&entity.Subscription{})

	return &SubscriptionRepo{
		db: db,
	}
}

// Create implements entity.SubscriptionRepo.
func (sr *SubscriptionRepo) Create(newSubscription entity.Subscription) (*entity.Subscription, error) {
	err := sr.db.Create(&newSubscription).Error
	if err != nil {
		return nil, err
	}

	return &newSubscription, nil
}

// Delete implements entity.SubscriptionRepo.
func (sr *SubscriptionRepo) Delete(id int64) (*entity.Subscription, error) {
	panic("unimplemented")
}

// ExistEmail implements entity.SubscriptionRepo.
func (sr *SubscriptionRepo) ExistEmail(email string) (*int64, error) {
	data := entity.Subscription{}

	err := sr.db.Select("id").Take(&data, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	id := int64(data.ID)
	return &id, nil
}

// Update implements entity.SubscriptionRepo.
func (sr *SubscriptionRepo) Update(id int64, newSubscription entity.Subscription) (*entity.Subscription, error) {
	var (
		subscription entity.Subscription
	)

	err := sr.db.First(&subscription, id).Error
	if err != nil {
		return nil, err
	}

	subscription = newSubscription
	subscription.ID = uint(id)
	err = sr.db.Save(&subscription).Error
	if err != nil {
		return nil, err
	}

	return &subscription, nil
}
