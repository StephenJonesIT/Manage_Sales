package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/report/model"
)

type reportStorage interface {
	ListReport(
		ctx context.Context,
		paging *common.Paging, 
		keymores ...string,
	) ([]model.Report, error)
}

type reportBiz struct{
	storage reportStorage
}

func NewReportStorage(storage reportStorage) (*reportBiz){
	return &reportBiz{storage: storage}
}

func(biz *reportBiz) ListReport(
	ctx context.Context, 
	paging *common.Paging,
) ([]model.Report, error){
	data, err := biz.storage.ListReport(ctx,paging)
	
	if err != nil {
		return nil, err
	}
	return data, nil
} 