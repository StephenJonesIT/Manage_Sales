package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"
)

type GetBonsaiStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.BonsaiItem, error)
}

type getBonsaiBiz struct {
	storage GetBonsaiStorage
}

func NewGetBonsaiBiz(storage GetBonsaiStorage) *getBonsaiBiz {
	return &getBonsaiBiz{storage: storage}
}

func (biz *getBonsaiBiz) GetBonsaiById(ctx context.Context, id string) (*model.BonsaiItem, error) {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"masp": id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.EntityName, err)
	}

	return data, nil
}
