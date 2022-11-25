package storage

import (
	"context"
	"gotodolist20221123/module/member/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.Member) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
