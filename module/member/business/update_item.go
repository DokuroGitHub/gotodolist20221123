package business

import (
	"context"
	"gotodolist20221123/module/member/model"
)

type UpdateItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.Member, error)

	UpdateItem(
		ctx context.Context,
		condition map[string]interface{},
		dataUpdate *model.Member,
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
	dataUpdate *model.Member,
) error {
	// step 1: Find item by conditions
	oldItem, err := business.storage.FindItem(ctx, condition)

	if err != nil {
		return err
	}

	//
	if dataUpdate.CreatedAt != oldItem.CreatedAt {
		return model.ErrCannotUpdateCreatedAt
	}

	// Step 2: call to storage to update item
	if err := business.storage.UpdateItem(ctx, condition, dataUpdate); err != nil {
		return err
	}

	return nil
}
