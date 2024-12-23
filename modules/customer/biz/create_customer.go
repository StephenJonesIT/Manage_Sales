package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"
)

type CreateCustomerStorage interface {
	CreateItem(ctx context.Context, data *model.CustomerCreate) error
}

type createBCustomerBiz struct {
	storage CreateCustomerStorage
}

func NewCreatecustomerBiz(storage CreateCustomerStorage) *createBCustomerBiz {
	return &createBCustomerBiz{storage: storage}
}

func (biz *createBCustomerBiz) CreateNewCustomer(ctx context.Context, data *model.CustomerCreate) error {

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}