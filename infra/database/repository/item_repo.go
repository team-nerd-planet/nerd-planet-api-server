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

func NewItemRepo(db *database.Database) entity.ItemRepo {
	return &ItemRepo{
		db: db,
	}
}

// Count implements entity.ItemRepo.
func (clr *ItemRepo) CountView(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64) (int64, error) {
	var (
		count int64
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if company != nil {
		where = append(where, "feed_name LIKE ?")
		param = append(param, fmt.Sprintf("%%%s%%", *company))
	}

	if companySizes != nil {
		where = append(where, "company_size IN ?")
		param = append(param, *companySizes)
	}

	if jobTags != nil {
		where = append(where, "job_tags_id_arr && ?") // `&&`: overlap (have elements in common)
		param = append(param, getArrToString(*jobTags))
	}

	if skillTags != nil {
		where = append(where, "skill_tags_id_arr && ?") // `&&`: overlap (have elements in common)
		param = append(param, getArrToString(*skillTags))
	}

	err := clr.db.Model(&entity.ItemView{}).
		Where(strings.Join(where, " AND "), param...).
		Count(&count).Error
	if err != nil {
		return -1, err
	}

	return count, nil
}

// FindAll implements entity.ItemRepo.
func (clr *ItemRepo) FindAllView(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64, perPage int, page int) ([]entity.ItemView, error) {
	var (
		items []entity.ItemView
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if company != nil {
		where = append(where, "feed_name LIKE ?")
		param = append(param, fmt.Sprintf("%%%s%%", *company))
	}

	if companySizes != nil && len(*companySizes) > 0 {
		where = append(where, "company_size IN ?")
		param = append(param, *companySizes)
	}

	if jobTags != nil && len(*jobTags) > 0 {
		where = append(where, "job_tags_id_arr && ?") // `&&`: overlap (have elements in common)
		param = append(param, getArrToString(*jobTags))
	}

	if skillTags != nil && len(*skillTags) > 0 {
		where = append(where, "skill_tags_id_arr && ?") // `&&`: overlap (have elements in common)
		param = append(param, getArrToString(*skillTags))
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
