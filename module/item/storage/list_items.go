package todostorage

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"
)

func (s *mysqlStorage) ListItem(
	ctx context.Context,
	condition map[string]interface{},
	paging *todoitemmodel.DataPaging,
) ([]todoitemmodel.ToDoItem, error) {
	offset := (paging.Page - 1) * paging.Limit

	var result []todoitemmodel.ToDoItem

	if err := s.db.Table(todoitemmodel.ToDoItem{}.TableName()).
		Where(condition).
		Count(&paging.Total).
		Offset(offset).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
