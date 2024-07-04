package handler

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/team-nerd-planet/api-server/infra/router/util"
	"github.com/team-nerd-planet/api-server/internal/controller/rest"
	_ "github.com/team-nerd-planet/api-server/internal/controller/rest/dto/tag_dto"
)

// ListJobTags
//
// @Summary			List Job Tag
// @Description		list job tags
// @Tags			tag
// @Schemes			http
// @Accept			json
// @Produce			json
// @Success			200 {object} []tag_dto.FindAllJobTagRes
// @Failure			400 {object} util.HTTPError
// @Failure			404 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/tag/job [get]
func ListJobTags(c iris.Context, ctrl rest.TagController) {
	res, ok := ctrl.FindAllJobTag()
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}

// ListSkillTags
//
// @Summary			List Skill Tag
// @Description		list skill tags
// @Tags			tag
// @Schemes			http
// @Accept			json
// @Produce			json
// @Success			200 {object} []tag_dto.FindAllSkillTagRes
// @Failure			400 {object} util.HTTPError
// @Failure			404 {object} util.HTTPError
// @Failure			500 {object} util.HTTPError
// @Router			/v1/tag/skill [get]
func ListSkillTags(c iris.Context, ctrl rest.TagController) {
	res, ok := ctrl.FindAllSkillTag()
	if !ok {
		util.NewError(c, http.StatusInternalServerError)
		return
	}

	c.JSON(res)
}
