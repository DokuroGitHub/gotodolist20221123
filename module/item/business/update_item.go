package business

import (
	"context"
	"gotodolist20221123/module/item/model"
)

type UpdateItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.ToDoItem, error)

	UpdateItem(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *model.ToDoItem,
	) error
}

type updateItemBusiness struct {
	storage UpdateItemStorage
}

func NewUpdateItemBusiness(storage UpdateItemStorage) *updateItemBusiness {
	return &updateItemBusiness{storage: storage}
}

func (business *updateItemBusiness) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *model.ToDoItem,
) error {
	// step 1: Find item by conditions
	oldItem, err := business.storage.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	// just a demo in case we dont allow update Finished item
	if oldItem.Status == "Finished" {
		return model.ErrCannotUpdateFinishedItem
	}

	// Step 2: call to storage to update item
	if err := business.storage.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
