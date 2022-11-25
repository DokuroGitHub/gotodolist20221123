package storage

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

func (s *mysqlStorage) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	if err := s.db.
		Table(model.Wallet{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
