package biz

import (
	"context"
	"manage_sales/modules/suppliers/model"
)

type DeleteSupplierStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.SupplierItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteSupplierBiz struct {
	storage DeleteSupplierStorage
}

func NewDeleteSupplierBiz(storage DeleteSupplierStorage) *deleteSupplierBiz {
	return &deleteSupplierBiz{storage: storage}
}

func (biz *deleteSupplierBiz) DeleteSupplierById(ctx context.Context, id string) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"mancc": id})
	if err != nil {
		return err
	}

	if data.TrangThai != nil && *data.TrangThai == model.BonsaiStatusDeleted {
		return model.ErrItemDeleted
	}

	if err := biz.storage.DeleteItem(ctx, map[string]interface{}{"mancc": id}); err != nil {
		return err
	}
	return nil
}