package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/account/model"
)

type CreateAccountStorage interface {
	CreateAccount(ctx context.Context, data *model.AccountCreate) error
}

type bizCreateAccount struct {
	storage CreateAccountStorage
}

func NewAccount(storage CreateAccountStorage) *bizCreateAccount {
	return &bizCreateAccount{storage: storage}
}

func (biz *bizCreateAccount) CreateAccountBiz(ctx context.Context, data *model.AccountCreate) error {

	if err := biz.storage.CreateAccount(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(model.Entity, err)
	}
	
	return nil
}