package todoitembusiness

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"
)

type DeleteTodoItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*todoitemmodel.ToDoItem, error)

	DeleteItem(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteBusiness struct {
	store DeleteTodoItemStorage
}

func NewDeleteToDoItemBiz(store DeleteTodoItemStorage) *deleteBusiness {
	return &deleteBusiness{store: store}
}

func (biz *deleteBusiness) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	// step 1: Find item by conditions
	_, err := biz.store.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	// Step 2: call to storage to delete item
	if err := biz.store.DeleteItem(ctx, condition); err != nil {
		return err
	}

	return nil
}
