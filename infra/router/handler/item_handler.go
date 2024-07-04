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

	c.JSON(res)
}

// IncreaseViewCount
//
// @Summary			Increase View Count
// @Description		increase item's view count
// @Tags			item
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		request body item_dto.IncreaseViewCountReq true "조회 수 증가 요청 내용"
// @Success			200 {object} item_dto.IncreaseViewCountRes
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/item/view_increase [post]
func IncreaseViewCount(c iris.Context, ctrl rest.ItemController) {
	req, err := util.ValidateBody[item_dto.IncreaseViewCountReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.IncreaseViewCount(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}

// IncreaseLikeCount
//
// @Summary			Increase Like Count
// @Description		increase item's like count
// @Tags			item
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param   		request body item_dto.IncreaseLikeCountReq true "좋아요 수 증가 요청 내용"
// @Success			200 {object} item_dto.IncreaseLikeCountRes
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/item/like_increase [post]
func IncreaseLikeCount(c iris.Context, ctrl rest.ItemController) {
	req, err := util.ValidateBody[item_dto.IncreaseLikeCountReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.IncreaseLikeCount(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}

// FindNextItems
//
// @Summary			Find Next Items
// @Description		find next items
// @Tags			item
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param			excluded_ids	query	[]int64		false	"제외할 글 ID 목록"
// @Param			limit			query	int32		true	"글 목록 개수 제한"
// @Success			200 {object} []item_dto.FindNextRes
// @Failure			400 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/item/next [get]
func FindNextItems(c iris.Context, ctrl rest.ItemController) {
	req, err := util.ValidateQuery[item_dto.FindNextReq](c)
	if err != nil {
		util.NewError(c, http.StatusBadRequest, err)
		return
	}

	res, ok := ctrl.FindNextItems(*req)
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}
