package handler

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/team-nerd-planet/api-server/infra/router/util"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
	_ "github.com/team-nerd-planet/api-server/internal/controller/rest/dto"
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/item_dto"
	_ "github.com/team-nerd-planet/api-server/internal/entity"
)

// ListItems
//
// @Summary			List item
// @Description		list items
// @Tags			item
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param			company			query	string						false	"회사 이름 검색 키워드"
// @Param			company_size	query	[]entity.CompanySizeType	false	"회사 규모 (0:스타트업, 1:중소기업, 2:중견기업, 3:대기업, 4:외국계)"
// @Param			job_tags		query	[]int64						false	"관련 직무 DB ID 배열"
// @Param			skill_tags		query	[]int64						false	"관련 스킬 DB ID 배열"
// @Param			page			query	int							true	"페이지"
// @Success			200 {object} dto.Paginated[[]item_dto.FindAllItemRes]
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/item [get]
func ListItems(c iris.Context, ctrl rest.ItemController) {
	req, err := util.ValidateQuery[item_dto.FindAllItemReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.FindAllItem(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}
