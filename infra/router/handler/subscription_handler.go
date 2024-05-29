package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
func Apply(c *gin.Context, ctrl rest.SubscriptionController) {
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

	c.JSON(http.StatusOK, res)
}

func Approve(c *gin.Context, ctrl rest.SubscriptionController) {
	req, err := util.ValidateQuery[subscription_dto.ApproveReq](c)
	if err != nil {
		c.Redirect(http.StatusBadRequest, "https://www.nerdplanet.app")
		return
	}

	res, ok := ctrl.Approve(*req)
	if !ok {
		c.Redirect(http.StatusInternalServerError, "https://www.nerdplanet.app")
		return
	}

	if res.Ok {
		c.Redirect(http.StatusFound, "https://www.nerdplanet.app")
	} else {
		c.Redirect(http.StatusBadRequest, "https://www.nerdplanet.app")
	}
}
