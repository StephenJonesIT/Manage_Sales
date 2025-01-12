package business

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

type insertBillStorage interface {
	CreateBill(ctx context.Context,data *model.CreateHoaDon) error
}

type insertBillBiz struct{
	storage insertBillStorage
}

func NewCreateBill(storage insertBillStorage) (*insertBillBiz){
	return &insertBillBiz{storage: storage}
}

func(biz *insertBillBiz) CreateBill(ctx context.Context, data *model.CreateHoaDon) error{
	if err := biz.storage.CreateBill(ctx, data); err != nil{
		return common.ErrCannotCreateEntity(model.Entity,err)
	}
	return nil
}
