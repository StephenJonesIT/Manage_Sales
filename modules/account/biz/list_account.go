package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/account/model"
)

type listAccountStorage interface {
	ListItem(
		ctx context.Context,
		paging *common.Paging,
		morekeys ...string,
	) ([]model.AccountItem, error)
}

type listAccountBiz struct {
	storage listAccountStorage
}

func NewListAccountBiz(storage listAccountStorage) *listAccountBiz {
	return &listAccountBiz{storage: storage}
}

func (binz *listAccountBiz) ListAccountById(
	ctx context.Context,
	paging *common.Paging,
) ([]model.AccountItem, error) {

	data, err := binz.storage.ListItem(ctx, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}