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

func (ic ItemController) IncreaseViewCount(req item_dto.IncreaseViewCountReq) (*item_dto.IncreaseViewCountRes, bool) {
	count, ok := ic.itemUcase.IncreaseViewCount(req.Id)
	if !ok {
		return nil, false
	}

	return &item_dto.IncreaseViewCountRes{
		ItemViewCount: count,
	}, true
}

func (ic ItemController) IncreaseLikeCount(req item_dto.IncreaseLikeCountReq) (*item_dto.IncreaseLikeCountRes, bool) {
	count, ok := ic.itemUcase.IncreaseLikeCount(req.Id)
	if !ok {
		return nil, false
	}

	return &item_dto.IncreaseLikeCountRes{
		ItemLikeCount: count,
	}, true
}

func (ic ItemController) FindNextItems(req item_dto.FindNextReq) ([]item_dto.FindNextRes, bool) {
	items, ok := ic.itemUcase.FindNextViewItem(req.ExcludedIds, req.Limit)
	if !ok {
		return []item_dto.FindNextRes{}, false
	}

	return item_dto.NewFindNextRes(*items), true
}
