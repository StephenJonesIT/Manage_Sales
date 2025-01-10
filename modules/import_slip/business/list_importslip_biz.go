package business

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/import_slip/model"
)

type ListImportSlipStorage interface {
	ListImportSlip(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		morekeys ...string,
	)([]model.PhieuNhap,error)
}

type listImportSlipBiz struct{
	storage ListImportSlipStorage
}

func NewListImportSlip(storage ListImportSlipStorage) (*listImportSlipBiz){
	return &listImportSlipBiz{storage: storage}
}

func(biz *listImportSlipBiz) ListImportSlips(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.PhieuNhap, error){
	data, err := biz.storage.ListImportSlip(ctx,filter,paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
