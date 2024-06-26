package usecase

import (
	"github.com/google/wire"
	"github.com/team-nerd-planet/api-server/internal/usecase/feed"
	"github.com/team-nerd-planet/api-server/internal/usecase/item"
	"github.com/team-nerd-planet/api-server/internal/usecase/smtp"
	"github.com/team-nerd-planet/api-server/internal/usecase/subscription"
	"github.com/team-nerd-planet/api-server/internal/usecase/tag"
	"github.com/team-nerd-planet/api-server/internal/usecase/token"
)

var UsecaseSet = wire.NewSet(
	item.NewItemUsecase,
	tag.NewJobTagUsecase,
	tag.NewSkillTagUsecase,
	subscription.NewSubscriptionUsecase,
	feed.NewFeedUsecase,
	token.NewTokenUsecase,
	smtp.NewSmtpUsecase,
)
