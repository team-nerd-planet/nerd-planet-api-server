package repository

import (
	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type JobTagRepo struct {
	db *database.Database
}

func NewJobTagRepo(db *database.Database) entity.JobTagRepo {
	return &JobTagRepo{
		db: db,
	}
}

// FindAll implements entity.JobTagRepo.
func (jtr *JobTagRepo) FindAll() ([]entity.JobTag, error) {
	var jobTags []entity.JobTag
	err := jtr.db.Find(&jobTags).Error
	return jobTags, err
}
