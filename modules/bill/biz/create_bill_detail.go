package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

type DetailBillStorage interface {
	InsertDetailBill(ctx context.Context,data *model.ChiTietHoaDon) error
}

type bizDetailBill struct{
	storage DetailBillStorage
}
func InitCreateBillDetail(storage DetailBillStorage) *bizDetailBill{
	return &bizDetailBill{storage: storage}
}

func(biz *bizDetailBill) CreateDetailBill(ctx context.Context, data *model.ChiTietHoaDon) error{
	if err := biz.storage.InsertDetailBill(ctx, data); err != nil  {
		return common.ErrCannotCreateEntity(model.EntityDetail, err)
	}
	return nil
}