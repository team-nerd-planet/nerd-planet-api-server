package tag

import (
	"log/slog"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type SkillTagUsecase struct {
	skillTagRepo entity.SkillTagRepo
}

func NewSkillTagUsecase(skillTagRepo entity.SkillTagRepo) SkillTagUsecase {
	return SkillTagUsecase{
		skillTagRepo: skillTagRepo,
	}
}

func (stu SkillTagUsecase) FindAll() (*[]entity.SkillTag, bool) {
	skillTags, err := stu.skillTagRepo.FindAll()
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return &skillTags, true
}
