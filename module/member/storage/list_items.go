package storage

import (
	"context"
	"fmt"
	"gotodolist20221123/module/member/model"

	"gorm.io/gorm/clause"
)

func (s *mysqlStorage) ListItems(
	ctx context.Context,
	condition map[string]interface{},
	paging *model.DataPaging,
	order map[string]bool,
) ([]model.Member, error) {
	offset := (paging.Page - 1) * paging.Limit
	orders := []clause.OrderByColumn{}
	for key, value := range order {
		_item := clause.OrderByColumn{Column: clause.Column{Name: key}, Desc: value}
		orders = append(orders, _item)
	}
	fmt.Println(orders)
	var result []model.Member

	if err := s.db.Table(model.Member{}.TableName()).
		Where(condition).
		Limit(paging.Limit).
		Count(&paging.Total).
		Offset(offset).
		// Order("created_at DESC")
		// Order(clause.OrderByColumn{Column: clause.Column{Name: "created_at"}, Desc: true}).
		Order(clause.OrderBy{Columns: orders}).
		// Order(orders).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
