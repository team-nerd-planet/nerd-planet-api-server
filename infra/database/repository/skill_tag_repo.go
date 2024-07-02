package repository

import (
	"log/slog"

	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type SkillTagRepo struct {
	db database.Database
}

func NewSkillTagRepo(db database.Database) entity.SkillTagRepo {
	if err := db.AutoMigrate(&entity.SkillTag{}); err != nil {
		slog.Error("Auto migrate SkillTag Entity.")
		panic(err)
	}
	
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
