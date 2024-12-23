package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"
)

type UpdateEmployeeStorage interface {
	GetEmployee(ctx context.Context, cond map[string]interface{}) (*model.Employee,error)
	UpdateEmployee(ctx context.Context, cond map[string]interface{}, dataUpdate *model.EmployeeUpdate) error
}

type updateEmployeeBiz struct {
	storage UpdateEmployeeStorage
}

func NewUpdateEmployeeBiz(storage UpdateEmployeeStorage) *updateEmployeeBiz {
	return &updateEmployeeBiz{storage: storage}
}

func (biz *updateEmployeeBiz) UpdateEmployeeById(ctx context.Context, id string, dataUpdate *model.EmployeeUpdate) error {

	data, err := biz.storage.GetEmployee(ctx, map[string]interface{}{"manv": id})
	if err != nil {
		if err == common.ErrRecordNotFound {
			return common.ErrCannotGetEntity(model.Entity, err)
		}
		return common.ErrCannotUpdateEntity(model.Entity, err)
	}

	if data.TrangThai != nil && *data.TrangThai == model.EmployeeStatusDeleted {
		return common.ErrEntityDelete(model.Entity, model.ErrEmployeeDeleted)
	}

	if err := biz.storage.UpdateEmployee(ctx, map[string]interface{}{"manv": id}, dataUpdate); err != nil {
		return common.ErrCannotUpdateEntity(model.Entity, err)
	}
	return nil
}