package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"
)

type ListSupplierStorage interface {
	ListItem(
		ctx context.Context,
		paging *common.Paging,
		morekeys ...string,
	) ([]model.SupplierItem, error)
}

type listSupplierBiz struct {
	storage ListSupplierStorage
}

func NewListSupplierBiz(storage ListSupplierStorage) *listSupplierBiz {
	return &listSupplierBiz{storage: storage}
}

func (binz *listSupplierBiz) ListSupplierById(
	ctx context.Context,
	paging *common.Paging,
) ([]model.SupplierItem, error) {

	data, err := binz.storage.ListItem(ctx, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}