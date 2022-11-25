package storage

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

func (s *mysqlStorage) ListItems(
	ctx context.Context,
	condition map[string]interface{},
	paging *model.DataPaging,
) ([]model.Wallet, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []model.Wallet

	if err := s.db.Table(model.Wallet{}.TableName()).
		Where(condition).
		Limit(paging.Limit).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
