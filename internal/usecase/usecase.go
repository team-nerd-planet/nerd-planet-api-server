package usecase

import (
	"github.com/google/wire"
	"github.com/team-nerd-planet/api-server/internal/usecase/item"
	"github.com/team-nerd-planet/api-server/internal/usecase/tag"
)

var UsecaseSet = wire.NewSet(item.NewItemUsecase, tag.NewJobTagUsecase, tag.NewSkillTagUsecase)
