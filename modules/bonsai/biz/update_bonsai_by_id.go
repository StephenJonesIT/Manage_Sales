package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bonsai/model"
)

type UpdateBonsaiStorage interface {
	GetItem(ctx context.Context, cond map[string]interface{}) (*model.BonsaiItem, error)
	UpdateItem(ctx context.Context, cond map[string]interface{}, dataUpdate *model.BonsaiItemUpdate) error
}

type updateBonsaiBiz struct {
	storage UpdateBonsaiStorage
}

func NewUpdateItemBiz(storage UpdateBonsaiStorage) *updateBonsaiBiz {
	return &updateBonsaiBiz{storage: storage}
}

func (biz *updateBonsaiBiz) UpdateItemById(ctx context.Context, id string, dataUpdate *model.BonsaiItemUpdate) error {

	data, err := biz.storage.GetItem(ctx, map[string]interface{}{"masp": id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.EntityName, err)
		}
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}

	if data.TrangThai != nil && *data.TrangThai == model.BonsaiStatusNgung {
		return common.ErrEntityDelete(model.EntityName, model.ErrItemDeleted)
	}

	if err := biz.storage.UpdateItem(ctx, map[string]interface{}{"masp": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.EntityName, err)
	}
	return nil
}