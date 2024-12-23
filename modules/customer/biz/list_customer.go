package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"
)

type ListCustomerStorage interface {
	ListItem(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]model.Customer, error)
}

type listCustomerBiz struct {
	storage ListCustomerStorage
}

func NewListCustomerBiz(storage ListCustomerStorage) *listCustomerBiz {
	return &listCustomerBiz{storage: storage}
}

func (binz *listCustomerBiz) ListCustomer(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.Customer, error) {

	data, err := binz.storage.ListItem(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}
