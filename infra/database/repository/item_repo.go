package repository

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
)

type ItemRepo struct {
	db *database.Database
}

func NewMemo(db *database.Database) entity.ItemRepo {
	return &ItemRepo{
		db: db,
	}
}

// Count implements entity.ItemRepo.
func (clr *ItemRepo) CountView(feedCompanySizeType *[]entity.CompanySizeType, itemTagIDArr *[]int64) (int64, error) {
	var (
		count int64
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if feedCompanySizeType != nil {
		where = append(where, "company_size IN ?")
		param = append(param, *feedCompanySizeType)
	}

	if itemTagIDArr != nil {
		where = append(where, "tag_id_arr @> ?") // `@>` = Does left array contain right array?
		param = append(param, getArrToString(*itemTagIDArr))
	}

	err := clr.db.Model(&entity.ViewItem{}).
		Where(strings.Join(where, " AND "), param...).
		Count(&count).Error
	if err != nil {
		return -1, err
	}

	return count, nil
}

// FindAll implements entity.ItemRepo.
func (clr *ItemRepo) FindAllView(feedCompanySizeType *[]entity.CompanySizeType, itemTagIDArr *[]int64, perPage int, page int) ([]entity.ViewItem, error) {
	var (
		items []entity.ViewItem
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if feedCompanySizeType != nil {
		where = append(where, "company_size IN ?")
		param = append(param, *feedCompanySizeType)
	}

	if itemTagIDArr != nil {
		where = append(where, "tag_id_arr @> ?") // `@>` = Does left array contain right array?
		param = append(param, getArrToString(*itemTagIDArr))
	}

	err := clr.db.
		Where(strings.Join(where, " AND "), param...).
		Offset((page - 1) * perPage).
		Limit(perPage).
		Find(&items).Error

	return items, err
}

func getArrToString(arr []int64) string {
	strArr := make([]string, len(arr))
	for i, v := range arr {
		strArr[i] = strconv.FormatInt(v, 10)
	}

	return fmt.Sprintf("{%s}", strings.Join(strArr, ","))
}
