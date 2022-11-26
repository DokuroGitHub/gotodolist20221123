package storage

import (
	"context"
	"gotodolist20221123/module/member/model"
)

func (s *mysqlStorage) ListItems(
	ctx context.Context,
	condition map[string]interface{},
	paging *model.DataPaging,
	order string,
) ([]model.Member, error) {
	offset := (paging.Page - 1) * paging.Limit
	var result []model.Member

	if err := s.db.Table(model.Member{}.TableName()).
		Where(condition).
		Limit(paging.Limit).
		Count(&paging.Total).
		Offset(offset).
		Order(order).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
