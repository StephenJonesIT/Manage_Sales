package business

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

type deleteBillStorage interface {
	GetBillById(ctx context.Context, cond map[string]interface{}) (*model.HoaDon, error)
	DeleteBill(ctx context.Context, cond map[string]interface{}) error
}

type deleteBillBiz struct {
	storage deleteBillStorage
}

func NewDeleteBill(storage deleteBillStorage) *deleteBillBiz {
	return &deleteBillBiz{storage: storage}
}

func (biz *deleteBillBiz) DeleteBillBiz(ctx context.Context, id string) error{
	data, err := biz.storage.GetBillById(ctx, map[string]interface{}{"mahd": id})
	if err != nil {
		return err
	}

	if data.TrangThai !=nil && *data.TrangThai == model.Deleted {
		return common.ErrEntityDelete(model.Entity,err)
	}

	if err := biz.storage.DeleteBill(ctx, map[string]interface{}{"mahd":id}); err!=nil {
		return err
	}

	return nil
}
