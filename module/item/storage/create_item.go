package storage

import (
	"context"
	"gotodolist20221123/module/item/model"
)

func (s *mysqlStorage) CreateItem(ctx context.Context, data *model.ToDoItem) error {
	if err := s.db.Create(data).Error; err != nil {
		return err
	}

	return nil
}
