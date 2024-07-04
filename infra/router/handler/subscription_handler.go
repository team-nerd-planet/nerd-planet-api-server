package handler

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/team-nerd-planet/api-server/infra/router/util"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/subscription_dto"
)

// Apply
//
// @Summary			Apply subscription
// @Description		apply for subscription
// @Tags			subscription
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		request body subscription_dto.ApplyReq true "contents for applying for subscription."
// @Success			200 {object} subscription_dto.ApplyRes
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/subscription/apply [post]
func Apply(c iris.Context, ctrl rest.SubscriptionController) {
	req, err := util.ValidateBody[subscription_dto.ApplyReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.Apply(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}

// Approve
//
// @Summary			Approve subscription
// @Description		approve for subscription
// @Tags			subscription
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		request body subscription_dto.ApproveReq true "contents for approving for subscription."
// @Success			200 {object} subscription_dto.ApproveRes
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/subscription/approve [post]
func Approve(c iris.Context, ctrl rest.SubscriptionController) {
	req, err := util.ValidateBody[subscription_dto.ApproveReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.Approve(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}
