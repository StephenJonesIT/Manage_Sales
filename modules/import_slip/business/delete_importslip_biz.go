package business

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/import_slip/model"
)

type deleteImportSlipStorage interface {
	GetImportSlipById(ctx context.Context, cond map[string]interface{}) (*model.PhieuNhap, error)
	DeleteImportSlip(ctx context.Context, cond map[string]interface{}) error
}

type deleteImportSlipBiz struct {
	storage deleteImportSlipStorage
}

func NewDeleteImportSlip(storage deleteImportSlipStorage) *deleteImportSlipBiz {
	return &deleteImportSlipBiz{storage: storage}
}

func (biz *deleteImportSlipBiz) DeleteImportSlipBiz(ctx context.Context, id string) error{
	data, err := biz.storage.GetImportSlipById(ctx, map[string]interface{}{"mapn": id})
	if err != nil {
		return err
	}

	if data.TrangThai !=nil && *data.TrangThai == model.Deleted {
		return common.ErrEntityDelete(model.EntitySlip,err)
	}

	if err := biz.storage.DeleteImportSlip(ctx, map[string]interface{}{"mapn":id}); err!=nil {
		return err
	}

	return nil
}
