package usecase

import (
	"github.com/google/wire"
	"github.com/team-nerd-planet/api-server/internal/usecase/item"
)

var UsecaseSet = wire.NewSet(item.NewItemUsecase)
