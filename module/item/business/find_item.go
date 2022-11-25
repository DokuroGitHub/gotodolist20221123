package business

import (
	"context"
	"gotodolist20221123/module/item/model"
)

type FindItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.ToDoItem, error)
}

type findItemBusiness struct {
	storage FindItemStorage
}

func NewFindItemBusiness(store FindItemStorage) *findItemBusiness {
	return &findItemBusiness{storage: store}
}

func (business *findItemBusiness) FindItem(ctx context.Context, condition map[string]interface{}) (*model.ToDoItem, error) {
	itemData, err := business.storage.FindItem(ctx, condition)

	if err != nil {
		return nil, err
	}

	return itemData, nil
}
