package storage

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

func (s *mysqlStorage) ListItems(
	ctx context.Context,
	where map[string]interface{},
	not map[string]interface{},
	or map[string]interface{},
	paging *model.DataPaging,
	order string,
) ([]model.Wallet, error) {
	offset := (paging.Page - 1) * paging.Limit
	var result []model.Wallet

	if err := s.db.Table(model.Wallet{}.TableName()).
		Where(where).
		Not(not).
		Or(or).
		Limit(paging.Limit).
		Count(&paging.Total).
		Offset(offset).
		Order(order).
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
