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
	res := make([]tag_dto.FindAllJobTagRes, 0)

	jobTags, ok := tc.jobTagUcase.FindAll()
	if !ok {
		return res, false
	}

	return tag_dto.NewFindAllJobTagRes(*jobTags), true
}

func (tc TagController) FindAllSkillTag() ([]tag_dto.FindAllSkillTagRes, bool) {
	res := make([]tag_dto.FindAllSkillTagRes, 0)

	skillTags, ok := tc.skillTagUcase.FindAll()
	if !ok {
		return res, false
	}

	return tag_dto.NewFindAllSkillTagRes(*skillTags), true
}
