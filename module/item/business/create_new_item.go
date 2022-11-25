package todoitembusiness

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"
)

type CreateTodoItemStorage interface {
	CreateItem(ctx context.Context, data *todoitemmodel.ToDoItem) error
}

type createBusiness struct {
	store CreateTodoItemStorage
}

func NewCreateToDoItemBiz(store CreateTodoItemStorage) *createBusiness {
	return &createBusiness{store: store}
}

func (biz *createBusiness) CreateNewItem(ctx context.Context, data *todoitemmodel.ToDoItem) error {
	if err := data.Validate(); err != nil {
		return err
	}

	// do not allow "finished" status when creating a new task
	data.Status = "Doing" // set to default

	if err := biz.store.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
