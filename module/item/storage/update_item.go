package storage

import (
	"context"
	"gotodolist20221123/module/item/model"
)

func (s *mysqlStorage) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *model.ToDoItem,
) error {
	if err := s.db.Where(condition).Updates(dataUpdate).Error; err != nil {
		return err
	}

	return nil
}
