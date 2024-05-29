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

func (iu ItemUsecase) CountViewItem(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64) (*int64, bool) {
	count, err := iu.itemRepo.CountView(company, companySizes, jobTags, skillTags)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return &count, true
}

func (mu ItemUsecase) FindAllViewItem(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64, perPage int, page int) (*[]entity.ItemView, bool) {
	viewItems, err := mu.itemRepo.FindAllView(company, companySizes, jobTags, skillTags, perPage, page)
	if err != nil {
		slog.Error(err.Error(), "error", err)
		return nil, false
	}

	return &viewItems, true
}
