package tag_dto

import "github.com/team-nerd-planet/api-server/internal/entity"

type FindAllJobTagRes struct {
	ID   uint   `json:"id"`   // 직무 DB ID
	Name string `json:"name"` // 직무 이름
}

func NewFindAllJobTagRes(jabTags []entity.JobTag) []FindAllJobTagRes {
	res := make([]FindAllJobTagRes, len(jabTags))

	for i, jabTag := range jabTags {
		res[i] = FindAllJobTagRes{
			ID:   jabTag.ID,
			Name: jabTag.Name,
		}
	}

	return res
}
