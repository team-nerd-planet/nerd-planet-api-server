package item

import (
	"log/slog"

	"github.com/team-nerd-planet/api-server/internal/entity"
)

type ItemUsecase struct {
	itemRepo entity.ItemRepo
}

func NewItemUsecase(itemRepo entity.ItemRepo) ItemUsecase {
	return ItemUsecase{
		itemRepo: itemRepo,
	}
}

func (iu ItemUsecase) CountViewItem(companySizeArr *[]entity.CompanySizeType, itemTagIDArr *[]int64) (*int64, bool) {
	count, err := iu.itemRepo.CountView(companySizeArr, itemTagIDArr)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &count, true
}

func (mu ItemUsecase) FindAllViewItem(companySizeArr *[]entity.CompanySizeType, itemTagIDArr *[]int64, perPage, page int) (*[]entity.ViewItem, bool) {
	viewItems, err := mu.itemRepo.FindAllView(companySizeArr, itemTagIDArr, perPage, page)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &viewItems, true
}
