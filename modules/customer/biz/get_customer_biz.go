package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/customer/model"
)

type GetCustomerStorage interface {
	GetCustomer(ctx context.Context, cond map[string] interface{}) (*model.Customer, error)
}

type getCustomerBiz struct{
	storage GetCustomerStorage
}

func NewGetCustomerBiz(storage GetCustomerStorage) *getCustomerBiz{
	return &getCustomerBiz{storage: storage}
}
func (biz *getCustomerBiz) GetCustomerById(ctx context.Context, id string) (*model.Customer, error){
	data, err := biz.storage.GetCustomer(ctx, map[string]interface{}{"makh":id})
	
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
} 

