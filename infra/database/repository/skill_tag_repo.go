package repository

import (
	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type SkillTagRepo struct {
	db database.Database
}

func NewSkillTagRepo(db database.Database) entity.SkillTagRepo {
	return &SkillTagRepo{
		db: db,
	}
}

// FindAll implements entity.SkillTagRepo.
func (str *SkillTagRepo) FindAll() ([]entity.SkillTag, error) {
	var skillTags []entity.SkillTag
	err := str.db.Find(&skillTags).Error
	return skillTags, err
}
