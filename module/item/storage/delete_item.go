package todostorage

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"
)

func (s *mysqlStorage) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {

	if err := s.db.
		Table(todoitemmodel.ToDoItem{}.TableName()).
		Where(condition).Delete(nil).Error; err != nil {
		return err
	}

	return nil
}
