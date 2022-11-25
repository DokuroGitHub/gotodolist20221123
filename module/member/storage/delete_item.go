package storage

import (
	"context"
	"gotodolist20221123/module/member/model"
)

func (s *mysqlStorage) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	if err := s.db.
		Table(model.Member{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
