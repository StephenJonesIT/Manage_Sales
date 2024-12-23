package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"
)

type UpdateCustomerStorage interface {
	GetCustomer(ctx context.Context, cond map[string]interface{}) (*model.Customer, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.CustomerUpdate) error
}

type updateCustomerBiz struct {
	storage UpdateCustomerStorage
}

func NewUpdateCustomerBiz(storage UpdateCustomerStorage) *updateCustomerBiz {
	return &updateCustomerBiz{storage: storage}
}

func (biz *updateCustomerBiz) UpdateCustomerById(ctx context.Context, id string, dataUpdate *model.CustomerUpdate) error {

	_, err := biz.storage.GetCustomer(ctx, map[string]interface{}{"makh": id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"makh": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}