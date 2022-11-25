package business

import (
	"context"
	"gotodolist20221123/module/item/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.ToDoItem) error
}

type createItemBusiness struct {
	storage CreateItemStorage
}

func NewCreateItemBusiness(storage CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{storage: storage}
}

func (business *createItemBusiness) CreateItem(ctx context.Context, data *model.ToDoItem) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := business.storage.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
