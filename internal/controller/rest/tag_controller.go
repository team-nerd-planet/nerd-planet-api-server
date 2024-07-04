package rest

import (
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/tag_dto"
	"github.com/team-nerd-planet/api-server/internal/usecase/tag"
)

type TagController struct {
	jobTagUcase   tag.JobTagUsecase
	skillTagUcase tag.SkillTagUsecase
}

func NewTagController(jobTagUsecase tag.JobTagUsecase, skillTagUcase tag.SkillTagUsecase) TagController {
	return TagController{
		jobTagUcase:   jobTagUsecase,
		skillTagUcase: skillTagUcase,
	}
}

func (tc TagController) FindAllJobTag() ([]tag_dto.FindAllJobTagRes, bool) {
	jobTags, ok := tc.jobTagUcase.FindAll()
	if !ok {
		return []tag_dto.FindAllJobTagRes{}, false
	}

	return tag_dto.NewFindAllJobTagRes(*jobTags), true
}

func (tc TagController) FindAllSkillTag() ([]tag_dto.FindAllSkillTagRes, bool) {
	skillTags, ok := tc.skillTagUcase.FindAll()
	if !ok {
		return []tag_dto.FindAllSkillTagRes{}, false
	}

	return tag_dto.NewFindAllSkillTagRes(*skillTags), true
}
