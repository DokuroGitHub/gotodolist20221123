package todoitembusiness

import (
	"context"
	todoitemmodel "gotodolist20221123/module/item/model"
)

type ListTodoItemStorage interface {
	ListItem(
		ctx context.Context,
		condition map[string]interface{},
		paging *todoitemmodel.DataPaging,
	) ([]todoitemmodel.ToDoItem, error)
}

type listBiz struct {
	store ListTodoItemStorage
}

func NewListToDoItemBiz(store ListTodoItemStorage) *listBiz {
	return &listBiz{store: store}
}

func (biz *listBiz) ListItems(ctx context.Context,
	condition map[string]interface{},
	paging *todoitemmodel.DataPaging,
) ([]todoitemmodel.ToDoItem, error) {
	result, err := biz.store.ListItem(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
