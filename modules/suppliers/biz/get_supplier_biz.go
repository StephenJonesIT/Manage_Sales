package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"
)

type GetSupplierStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.SupplierItem, error)
}

type getSupplierBiz struct {
	storage GetSupplierStorage
}

func NewGetSupplierBiz(storage GetSupplierStorage) *getSupplierBiz {
	return &getSupplierBiz{storage: storage}
}

func (biz *getSupplierBiz) GetSupplierById(ctx context.Context, id string) (*model.SupplierItem, error) {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"mancc": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}