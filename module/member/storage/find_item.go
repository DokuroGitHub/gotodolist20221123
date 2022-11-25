package storage

import (
	"context"
	"gotodolist20221123/module/member/model"

	"gorm.io/gorm"
)

func (s *mysqlStorage) FindItem(
	ctx context.Context,
	condition map[string]interface{},
) (*model.Member, error) {
	var itemData model.Member

	if err := s.db.Where(condition).First(&itemData).Error; err != nil {
		if err == gorm.ErrRecordNotFound { // data not found
			return nil, model.ErrItemNotFound
		}

		return nil, err // other errors
	}

	return &itemData, nil
}
