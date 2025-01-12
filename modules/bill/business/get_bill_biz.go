package business

import (
	"context"
	"manage_sales/modules/bill/model"
)

type getBillStorage interface {
	GetBillById(ctx context.Context, cond map[string] interface{})(*model.HoaDon, error)
}

type getBillBiz struct{
	storage getBillStorage
}

func NewGetBill(storage getBillStorage) (*getBillBiz){
	return &getBillBiz{storage: storage}
}

func(biz *getBillBiz) GetBill(ctx context.Context, id string) (*model.HoaDon,error){
	
	result, err := biz.storage.GetBillById(ctx, map[string]interface{}{"mahd":id})
	if err != nil {
		return nil, err
	}
	return result, nil
}