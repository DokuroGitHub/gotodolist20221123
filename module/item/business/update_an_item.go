package todoitembusiness

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"
)

type UpdateTodoItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*todoitemmodel.ToDoItem, error)

	UpdateItem(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *todoitemmodel.ToDoItem,
	) error
}

type updateBiz struct {
	store UpdateTodoItemStorage
}

func NewUpdateToDoItemBiz(store UpdateTodoItemStorage) *updateBiz {
	return &updateBiz{store: store}
}

func (biz *updateBiz) UpdateItem(
	ctx context.Context,
	condition map[string]interface{},
	dataUpdate *todoitemmodel.ToDoItem,
) error {
	// step 1: Find item by conditions
	oldItem, err := biz.store.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	// just a demo in case we dont allow update Finished item
	if oldItem.Status == "Finished" {
		return todoitemmodel.ErrCannotUpdateFinishedItem
	}

	// Step 2: call to storage to update item
	if err := biz.store.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
