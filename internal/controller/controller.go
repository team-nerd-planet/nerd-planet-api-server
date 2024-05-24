package controller

import (
	"github.com/google/wire"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
)

var ControllerSet = wire.NewSet(rest.NewItemController, rest.NewTagController)
