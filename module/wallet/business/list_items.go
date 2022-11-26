package business

import (
	"context"
	"gotodolist20221123/module/wallet/model"
)

type ListItemsStorage interface {
	ListItems(
		ctx context.Context,
		where map[string]interface{},
		not map[string]interface{},
		or map[string]interface{},
		paging *model.DataPaging,
		order string,
	) ([]model.Wallet, error)
}

type listItemsBusiness struct {
	storage ListItemsStorage
}

func NewListItemsBusiness(storage ListItemsStorage) *listItemsBusiness {
	return &listItemsBusiness{storage: storage}
}

func (business *listItemsBusiness) ListItems(
	ctx context.Context,
	where map[string]interface{},
	not map[string]interface{},
	or map[string]interface{},
	paging *model.DataPaging,
	order string,
) ([]model.Wallet, error) {
	result, err := business.storage.ListItems(ctx, where, not, or, paging, order)

	if err != nil {
		return nil, err
	}

	return result, err
}
