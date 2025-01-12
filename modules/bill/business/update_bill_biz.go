package business

import (
	"context"
	"errors"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

type updateBillStorage interface {
	GetBillById(ctx context.Context, cond map[string]interface{}) (*model.HoaDon, error)
	UpdateBillById(ctx context.Context, cond map[string]interface{},updateData *model.UpdateHoaDon) (error)
}

type updateBillBiz struct{
	storage updateBillStorage
}

func NewUpdateBill(storage updateBillStorage) (*updateBillBiz){
	return &updateBillBiz{storage: storage}
}

func(biz *updateBillBiz) UpdateBill(ctx context.Context, id string, updateData *model.UpdateHoaDon) error{
	data, err := biz.storage.GetBillById(ctx,map[string]interface{}{"mahd":id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.Entity,err)
		}
		return common.ErrCannotUpdateEntity(model.Entity, err)
	}

	if data.TrangThai != nil && *data.TrangThai==model.Deleted{
		return common.ErrEntityDelete(model.Entity,errors.New("Bill deleted"))
	}

	if err := biz.storage.UpdateBillById(ctx, map[string]interface{}{"mahd":id},updateData); err!=nil {
		return common.ErrCannotUpdateEntity(model.Entity,err)
	}

	return nil
}