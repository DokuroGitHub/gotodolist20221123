package business

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

type FindItemStorage interface {
	FindItem(
		ctx context.Context,
		condition map[string]interface{},
	) (*model.Wallet, error)
}

type findItemBusiness struct {
	storage FindItemStorage
}

func NewFindItemBusiness(store FindItemStorage) *findItemBusiness {
	return &findItemBusiness{storage: store}
}

func (biz *findItemBusiness) FindItem(ctx context.Context, condition map[string]interface{}) (*model.Wallet, error) {
	itemData, err := biz.storage.FindItem(ctx, condition)

	if err != nil {
		return nil, err
	}

	return itemData, nil
}
