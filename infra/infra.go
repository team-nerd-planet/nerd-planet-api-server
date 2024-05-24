package infra

import (
	"github.com/google/wire"
	"github.com/team-nerd-planet/api-server/infra/config"
	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/infra/database/repository"
	"github.com/team-nerd-planet/api-server/infra/router"
)

var InfraSet = wire.NewSet(
	config.NewConfig,
	router.NewRouter,
	database.NewDatabase,
	repository.NewItemRepo,
	repository.NewJobTagRepo,
	repository.NewSkillTagRepo,
)
