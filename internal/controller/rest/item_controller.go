package rest

import (
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto"
	"github.com/team-nerd-planet/api-server/internal/controller/rest/dto/item_dto"
	"github.com/team-nerd-planet/api-server/internal/usecase/item"
)

type ItemController struct {
	itemUcase item.ItemUsecase
}

func NewItemController(itemUsecase item.ItemUsecase) ItemController {
	return ItemController{
		itemUcase: itemUsecase,
	}
}

func (ic ItemController) FindAllItem(req item_dto.FindAllItemReq) (dto.Paginated[[]item_dto.FindAllItemRes], bool) {
	const perPage = 100
	data := dto.Paginated[[]item_dto.FindAllItemRes]{}

	totalCount, ok := ic.itemUcase.CountViewItem(req.Company, req.CompanySize, req.JobTags, req.SkillTags)
	if !ok {
		return data, false
	}

	items, ok := ic.itemUcase.FindAllViewItem(req.Company, req.CompanySize, req.JobTags, req.SkillTags, perPage, *req.Page)
	if !ok {
		return data, false
	}

	itemRes := make([]item_dto.FindAllItemRes, len(*items))
	for i, item := range *items {
		itemRes[i] = item_dto.NewFindAllItemRes(item)
	}

	data = dto.NewPaginatedRes(itemRes, *req.Page, perPage, *totalCount)
	return data, true
}
