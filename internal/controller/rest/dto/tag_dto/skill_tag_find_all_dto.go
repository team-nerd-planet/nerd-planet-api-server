package tag_dto

import "github.com/team-nerd-planet/api-server/internal/entity"

type FindAllSkillTagRes struct {
	ID   uint   `json:"id"`   // 스킬 DB ID
	Name string `json:"name"` // 스킬 이름
}

func NewFindAllSkillTagRes(skillTags []entity.SkillTag) []FindAllSkillTagRes {
	res := make([]FindAllSkillTagRes, len(skillTags))

	for i, skillTag := range skillTags {
		res[i] = FindAllSkillTagRes{
			ID:   skillTag.ID,
			Name: skillTag.Name,
		}
	}

	return res
}
