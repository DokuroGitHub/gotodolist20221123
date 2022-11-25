package storage

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.Wallet) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
