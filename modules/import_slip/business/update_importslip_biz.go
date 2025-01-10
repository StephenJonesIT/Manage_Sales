package business

import (
	"context"
	"errors"
	"manage_sales/common"
	"manage_sales/modules/import_slip/model"
)

type updateImportSlipStorage interface {
	GetImportSlipById(ctx context.Context, cond map[string]interface{}) (*model.PhieuNhap, error)
	UpdateImportSlipById(ctx context.Context, cond map[string]interface{},updateData *model.PhieuNhapUpdate) (error)
}

type updateImportSlipBiz struct{
	storage updateImportSlipStorage
}

func NewUpdateImportSlip(storage updateImportSlipStorage) (*updateImportSlipBiz){
	return &updateImportSlipBiz{storage: storage}
}

func(biz *updateImportSlipBiz) UpdateImportSlip(ctx context.Context, id string, updateData *model.PhieuNhapUpdate) error{
	data, err := biz.storage.GetImportSlipById(ctx,map[string]interface{}{"mapn":id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.EntitySlip,err)
		}
		return common.ErrCannotUpdateEntity(model.EntitySlip, err)
	}

	if data.TrangThai != nil && *data.TrangThai==model.Deleted{
		return common.ErrEntityDelete(model.EntitySlip,errors.New("Import Slip deleted"))
	}

	if err := biz.storage.UpdateImportSlipById(ctx, map[string]interface{}{"mapn":id},updateData); err!=nil {
		return common.ErrCannotUpdateEntity(model.EntitySlip,err)
	}

	return nil
}