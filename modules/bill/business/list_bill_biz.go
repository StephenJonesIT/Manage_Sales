package business

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

type listBillStorage interface {
	ListBill(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		morekeys ...string,
	)([]model.HoaDon,error)
}

type listBillBiz struct{
	storage listBillStorage
}

func NewListBills(storage listBillStorage) (*listBillBiz){
	return &listBillBiz{storage: storage}
}

func(biz *listBillBiz) ListBills(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.HoaDon, error){
	data, err := biz.storage.ListBill(ctx,filter,paging)
	if err != nil {
		return nil, err
	}
	return data, nil
}
