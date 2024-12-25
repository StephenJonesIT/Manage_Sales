package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"

)

type CreateSupplierStorage interface {
	CreateItem(ctx context.Context, data *model.SupplierItemCreate) error
}

type createSupplierBiz struct {
	storage CreateSupplierStorage
}

func NewCreateSupplierBiz(storage CreateSupplierStorage) *createSupplierBiz{
	return &createSupplierBiz{storage: storage}
}

func (biz *createSupplierBiz) CreateNewSupplier(ctx context.Context, data *model.SupplierItemCreate) error {

	// title := strings.TrimSpace(data.Ten)

	// if title != "" {
	// 	return model.ErrTitleIsBlank
	// }

	if err := biz.storage.CreateItem(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.EntityName, err)
	}

	return nil
}