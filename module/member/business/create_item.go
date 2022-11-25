package business

import (
	"context"
	"gotodolist20221123/module/member/model"
)

type CreateItemStorage interface {
	CreateItem(ctx context.Context, data *model.Member) error
}

type createItemBusiness struct {
	storage CreateItemStorage
}

func NewCreateItemBusiness(store CreateItemStorage) *createItemBusiness {
	return &createItemBusiness{storage: store}
}

func (business *createItemBusiness) CreateNewItem(ctx context.Context, data *model.Member) error {
	if err := data.Validate(); err != nil {
		return err
	}

	if err := business.storage.CreateItem(ctx, data); err != nil {
		return err
	}

	return nil
}
