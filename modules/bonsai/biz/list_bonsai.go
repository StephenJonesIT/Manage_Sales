package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"
)

type ListBonsaiStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]model.BonsaiItem, error)
}

type listBonsaiBiz struct {
	storage ListBonsaiStorage
}

func NewListBonsaiBiz(storage ListBonsaiStorage) *listBonsaiBiz {
	return &listBonsaiBiz{storage: storage}
}

func (binz *listBonsaiBiz) ListItemById(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.BonsaiItem, error) {

	data, err := binz.storage.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}