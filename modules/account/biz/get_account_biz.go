package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/account/model"
)

type StorageAccount interface {
	FindAccount(context context.Context, cond map[string]interface{}) (*model.AccountItem, error)
}

type bizGetAccount struct{
	storage StorageAccount
}

func NewGetAccountBiz(storage StorageAccount) *bizGetAccount{
	return &bizGetAccount{storage: storage}
}

func (biz *bizGetAccount) FindAccountById(ctx context.Context, id string) (*model.AccountItem, error){
	data, err := biz.storage.FindAccount(ctx, map[string]interface{}{"ma_tai_khoan": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.Entity, err)
	}
	return data, nil
}

func (biz *bizGetAccount) FindAccountByUsername(ctx context.Context, username string) (*model.AccountItem, error){
	data, err := biz.storage.FindAccount(ctx, map[string]interface{}{"ten_dang_nhap": username})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.Entity, err)
	}
	return data, nil
}