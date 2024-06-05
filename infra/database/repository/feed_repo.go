package repository

import (
	"fmt"
	"strings"

	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type FeedRepo struct {
	db database.Database
}

func NewFeedRepo(db database.Database) entity.FeedRepo {
	return &FeedRepo{
		db: db,
	}
}

// FindAll implements entity.FeedRepo.
func (fr *FeedRepo) FindAll(keyword *string) ([]entity.Feed, error) {
	var (
		feeds []entity.Feed
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if keyword != nil {
		where = append(where, "name LIKE ?")
		param = append(param, fmt.Sprintf("%s%%", *keyword))
	}

	err := fr.db.
		Where(strings.Join(where, " AND "), param...).
		Order("name").
		Find(&feeds).Error

	return feeds, err
}
