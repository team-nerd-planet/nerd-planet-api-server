package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/item_dto"
	_ "github.com/team-nerd-planet/api-server/internal/entity"
)

// ListItems
//
// @Summary			List item
// @Description		items
// @Tags			item
// @Schemes			http
// @Accept			json
// @Produce			json
// @Param			company_size	query	[]entity.CompanySizeType	false	"search by company_size"		collectionFormat(multi)
// @Param			tags			query	[]int64						false	"search by tag"					collectionFormat(multi)
// @Param			page			query	int							true	"content search by keyword"		minimum(1)
// @Success			200 {object} item_dto.FindAllItemRes
// @Failure			400
// @Failure			404
// @Failure			500
// @Router			/v1/item [get]
func ListItems(c *gin.Context, ctrl rest.ItemController) {
	req, ok := validateQuery[item_dto.FindAllItemReq](c)
	if !ok {
		c.Status(http.StatusBadRequest)
		return
	}

	res, ok := ctrl.FindAllItem(*req)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, res)
}

func validateBody[T any](c *gin.Context) (*T, bool) {
	var input T

	if err := c.ShouldBind(&input); err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &input, true
}

func validateQuery[T any](c *gin.Context) (*T, bool) {
	var input T

	if err := c.ShouldBindQuery(&input); err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &input, true
}

func validateInt64Param(c *gin.Context, key string) (*int64, bool) {
	param := c.Param(key)
	id, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &id, true
}
