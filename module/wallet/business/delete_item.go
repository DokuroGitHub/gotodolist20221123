package business

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

type DeleteItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.Wallet, error)

	DeleteItem(
		ctx context.Context,
		condition map[string]interface{},
	) error
}

type deleteItemBusiness struct {
	storage DeleteItemStorage
}

func NewDeleteBusiness(storage DeleteItemStorage) *deleteItemBusiness {
	return &deleteItemBusiness{storage: storage}
}

func (business *deleteItemBusiness) DeleteItem(
	ctx context.Context,
	condition map[string]interface{},
) error {
	// step 1: Find item by conditions
	_, err := business.storage.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	// Step 2: call to storage to delete item
	if err := business.storage.DeleteItem(ctx, condition); err != nil {
		return err
	}

	return nil
}
