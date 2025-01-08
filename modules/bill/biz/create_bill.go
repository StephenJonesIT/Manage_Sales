package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

type BillStorage interface {
	InsertBill(ctx context.Context,data *common.HoaDon) error
}

type bizBill struct{
	storage BillStorage
}
func InitCreateBill(storage BillStorage) *bizBill{
	return &bizBill{storage: storage}
}

func(biz *bizBill) CreateBill(ctx context.Context, data *common.HoaDon) error{
	if err := biz.storage.InsertBill(ctx, data); err != nil  {
		return common.ErrCannotCreateEntity(model.Entity, err)
	}
	return nil
}