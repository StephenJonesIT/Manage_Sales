package biz

import (
	"context"
)

type DeleteAccountStorage interface {
	DeleteAccountByID(ctx context.Context, cond map[string]interface{}) error
}

type deleteAccountBiz struct {
	storage DeleteAccountStorage
}

func NewDeleteAccountBiz(storage DeleteAccountStorage) *deleteAccountBiz {
	return &deleteAccountBiz{storage: storage}
}

func (biz *deleteAccountBiz) DeleteAccount(ctx context.Context, id string) error {

	if err := biz.storage.DeleteAccountByID(ctx, map[string]interface{}{"ma_tai_khoan": id}); err != nil {
		return err
	}
	return nil
}