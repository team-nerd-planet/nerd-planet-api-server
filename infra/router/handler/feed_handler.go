package handler

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/team-nerd-planet/api-server/infra/router/util"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/feed_dto"
)

// SearchFeedName
//
// @Summary			Search Feed Name
// @Description		search feed's name
// @Tags			feed
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param			name_keyword	query	string	false	"회사 이름 검색 키워드"
// @Success			200 {object} []feed_dto.SearchRes
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/feed/search [get]
func SearchFeedName(c iris.Context, ctrl rest.FeedController) {
	req, err := util.ValidateQuery[feed_dto.SearchReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.Search(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}
