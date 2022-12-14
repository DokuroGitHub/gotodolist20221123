package storage

import (
	"context"
	"gotodolist20221123/module/item/model"

	"gorm.io/gorm"
)

func (s *mysqlStorage) FindItem(
	ctx context.Context,
	condition map[string]interface{},
) (*model.ToDoItem, error) {
	var itemData model.ToDoItem

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, model.ErrItemNotFound
		}

		return nil, err // other errors
	}

	return &itemData, nil
}
