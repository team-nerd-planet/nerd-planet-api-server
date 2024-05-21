//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/team-nerd-planet/api-server/infra"
	"github.com/team-nerd-planet/api-server/infra/router"
	"github.com/team-nerd-planet/api-server/internal/controller"
	"github.com/team-nerd-planet/api-server/internal/usecase"
)

func InitServer() (router.Router, error) {
	wire.Build(infra.InfraSet, controller.ControllerSet, usecase.UsecaseSet)
	return router.Router{}, nil
}
