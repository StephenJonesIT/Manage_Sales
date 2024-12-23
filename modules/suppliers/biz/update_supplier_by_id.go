package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/suppliers/model"
)

type UpdateSupplierStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.SupplierItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.SupplierItemUpdate) error
}

type updateSupplierBiz struct {
	storage UpdateSupplierStorage
}

func NewUpdateSupplierBiz(storage UpdateSupplierStorage) *updateSupplierBiz {
	return &updateSupplierBiz{storage: storage}
}

func (biz *updateSupplierBiz) UpdateSupplierById(ctx context.Context, id string, dataUpdate *model.SupplierItemUpdate) error {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"mancc": id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if data.TrangThai != nil && *data.TrangThai == model.BonsaiStatusDeleted {
		return common.ErrEntityDelete(model.EntityName, model.ErrItemDeleted)
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"mancc": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}