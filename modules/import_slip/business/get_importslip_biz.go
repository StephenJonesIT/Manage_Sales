package business

import (
	"context"
	"manage_sales/modules/import_slip/model"
)

type getImportSlipStorage interface {
	GetImportSlipById(ctx context.Context, cond map[string] interface{})(*model.PhieuNhap, error)
}

type getImportSlipBiz struct{
	storage getImportSlipStorage
}

func NewGetImportSlip(storage getImportSlipStorage) (*getImportSlipBiz){
	return &getImportSlipBiz{storage: storage}
}

func(biz *getImportSlipBiz) GetImportSlip(ctx context.Context, id string) (*model.PhieuNhap,error){
	
	result, err := biz.storage.GetImportSlipById(ctx, map[string]interface{}{"mapn":id})
	if err != nil {
		return nil, err
	}
	return result, nil
}