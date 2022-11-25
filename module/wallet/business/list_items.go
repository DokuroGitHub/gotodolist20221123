package business

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

type ListItemsStorage interface {
	ListItems(
		ctx context.Context,
		condition map[string]interface{},
		paging *model.DataPaging,
	) ([]model.Wallet, error)
}

type listItemsBusiness struct {
	storage ListItemsStorage
}

func NewListItemsBusiness(storage ListItemsStorage) *listItemsBusiness {
	return &listItemsBusiness{storage: storage}
}

func (business *listItemsBusiness) ListItems(ctx context.Context,
	condition map[string]interface{},
	paging *model.DataPaging,
) ([]model.Wallet, error) {
	result, err := business.storage.ListItems(ctx, condition, paging)

	if err != nil {
		return nil, err
	}

	return result, err
}
