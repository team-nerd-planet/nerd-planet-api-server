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
		slog.Error(err.Error())
		return nil, false
	}

	return &count, true
}

func (iu ItemUsecase) FindAllViewItem(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64, perPage int, page int) (*[]entity.ItemView, bool) {
	viewItems, err := iu.itemRepo.FindAllView(company, companySizes, jobTags, skillTags, perPage, page)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &viewItems, true
}

func (iu ItemUsecase) IncreaseViewCount(id int64) (int64, bool) {
	count, err := iu.itemRepo.IncreaseViewCount(id)
	if err != nil {
		slog.Error(err.Error())
		return -1, false
	}

	return count, true
}

func (iu ItemUsecase) IncreaseLikeCount(id int64) (int64, bool) {
	count, err := iu.itemRepo.IncreaseLikeCount(id)
	if err != nil {
		slog.Error(err.Error())
		return -1, false
	}

	return count, true
}

func (iu ItemUsecase) FindNextViewItem(ids []int64, limit int32) (*[]entity.ItemView, bool) {
	viewItems, err := iu.itemRepo.FindAllViewByExcludedIds(ids, limit)
	if err != nil {
		slog.Error(err.Error())
		return nil, false
	}

	return &viewItems, true
}
