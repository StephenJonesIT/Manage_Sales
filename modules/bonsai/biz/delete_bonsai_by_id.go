package biz

import (
	"context"
	"manage_sales/modules/bonsai/model"
)

type DeleteBonsaiStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.BonsaiItem, error)
	DeleteItem(ctx context.Context, cond map[string]interface{}) error
}

type deleteBonsaiBiz struct {
	storage DeleteBonsaiStorage
}

func NewDeleteItemBiz(storage DeleteBonsaiStorage) *deleteBonsaiBiz {
	return &deleteBonsaiBiz{storage: storage}
}

func (biz *deleteBonsaiBiz) DeleteItemById(ctx context.Context, id string) error {
	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"masp": id})
	if err != nil {
		return err
	}

	if data.TrangThai != nil && *data.TrangThai == model.BonsaiStatusNgung {
		return model.ErrItemDeleted
	}

	if err := biz.storage.DeleteItem(ctx, map[string]interface{}{"masp": id}); err != nil {
		return err
	}
	return nil
}