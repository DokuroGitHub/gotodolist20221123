package todostorage

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"

	"gorm.io/gorm"
)

func (s *mysqlStorage) FindItem(
	ctx context.Context,
	condition map[string]interface{},
) (*todoitemmodel.ToDoItem, error) {
	var itemData todoitemmodel.ToDoItem

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, todoitemmodel.ErrItemNotFound
		}

		return nil, err // other errors
	}

	return &itemData, nil
}
