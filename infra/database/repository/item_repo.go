package repository

import (
	"errors"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"github.com/team-nerd-planet/api-server/infra/database"
	"github.com/team-nerd-planet/api-server/internal/entity"
	"gorm.io/gorm"
)

type ItemRepo struct {
	db database.Database
}

func NewItemRepo(db database.Database) entity.ItemRepo {
	if err := db.AutoMigrate(&entity.Item{}); err != nil {
		slog.Error("Auto migrate Item Entity.")
		panic(err)
	}

	db.Migrator().CreateView("vw_items", gorm.ViewOption{
		Replace: true,
		Query: db.
			Table("items i").
			Select(`
				i."id" as item_id, 
				i.title as item_title, 
				i.description as item_description, 
				i."link" as item_link,
				i.thumbnail as item_thumbnail,
				i.published as item_published,
				i.summary as item_summary,
				i.views as item_views,
				i.likes as item_likes,
				f."id" as feed_id,
				f."name" as feed_name, 
				f.title as feed_title, 
				f."link" as feed_link,
				f.company_size as company_size, 
				job_tags.id_arr as job_tags_id_arr,
				skill_tags.id_arr as skill_tags_id_arr
			`).
			Joins(`LEFT JOIN feeds f ON f."id" = i.feed_id`).
			Joins(`LEFT JOIN (?) as job_tags ON job_tags.item_id = i."id"`,
				db.Table("item_job_tags").
					Select("item_id, array_agg(job_tag_id) as id_arr").
					Group("item_id")).
			Joins(`LEFT JOIN (?) as skill_tags ON skill_tags.item_id = i."id"`,
				db.Table("item_skill_tags").
					Select("item_id, array_agg(skill_tag_id) as id_arr").
					Group("item_id")).
			Order(`i.published desc, i."id" desc`)})

	return &ItemRepo{
		db: db,
	}
}

// Count implements entity.ItemRepo.
func (ir *ItemRepo) CountView(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64) (int64, error) {
	var (
		count int64
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if company != nil {
		where = append(where, "feed_name LIKE ?")
		param = append(param, fmt.Sprintf("%s%%", *company))
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

	err := ir.db.Model(&entity.ItemView{}).
		Where(strings.Join(where, " AND "), param...).
		Count(&count).Error
	if err != nil {
		return -1, err
	}

	return count, nil
}

// FindAll implements entity.ItemRepo.
func (ir *ItemRepo) FindAllView(company *string, companySizes *[]entity.CompanySizeType, jobTags, skillTags *[]int64, perPage int, page int) ([]entity.ItemView, error) {
	var (
		items []entity.ItemView
		where = make([]string, 0)
		param = make([]interface{}, 0)
	)

	if company != nil {
		where = append(where, "feed_name LIKE ?")
		param = append(param, fmt.Sprintf("%s%%", *company))
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

	err := ir.db.
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
