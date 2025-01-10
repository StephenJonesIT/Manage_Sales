package business

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/import_slip/model"
)

type insertImportSlipStorage interface {
	CreateImportSlip(ctx context.Context,data *model.PhieuNhapCreate) error
}

type insertImportSlipBiz struct{
	storage insertImportSlipStorage
}

func NewCreateImportSlip(storage insertImportSlipStorage) (*insertImportSlipBiz){
	return &insertImportSlipBiz{storage: storage}
}

func(biz *insertImportSlipBiz) CreateImportSlip(ctx context.Context, data *model.PhieuNhapCreate) error{
	if err := biz.storage.CreateImportSlip(ctx, data); err != nil{
		return common.ErrCannotCreateEntity(model.EntitySlip,err)
	}
	return nil
}
