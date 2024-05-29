package rest

import (
	"log/slog"

	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/subscription_dto"
	"github.com/team-nerd-planet/api-server/internal/usecase/subscription"
)

type SubscriptionController struct {
	subscriptionUcase subscription.SubscriptionUsecase
}

func NewSubscriptionController(subscriptionUcase subscription.SubscriptionUsecase) SubscriptionController {
	return SubscriptionController{
		subscriptionUcase: subscriptionUcase,
	}
}

func (sc SubscriptionController) Apply(req subscription_dto.ApplyReq) (*subscription_dto.ApplyRes, bool) {
	slog.Info("SubscriptionController_Apply", "req", req)

	subscription, ok := sc.subscriptionUcase.Apply(req.NewSubscription())
	if !ok {
		return nil, false
	}

	slog.Info("Result", "subscription", subscription)

	result := false
	if subscription != nil {
		result = true
	}

	return &subscription_dto.ApplyRes{
		Ok: result,
	}, true
}

func (sc SubscriptionController) Approve(req subscription_dto.ApproveReq) (*subscription_dto.ApproveRes, bool) {
	subscription, ok := sc.subscriptionUcase.Subscribe(req.Token)
	if !ok {
		return nil, false
	}

	result := false
	if subscription != nil {
		result = true
	}

	return &subscription_dto.ApproveRes{
		Ok: result,
	}, true
}
