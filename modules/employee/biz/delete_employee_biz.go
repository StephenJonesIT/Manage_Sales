package biz

import (
	"context"
	"manage_sales/modules/employee/model"
)

type DeleteEmployeeStorage interface {
	GetEmployee(ctx context.Context, cond map[string]interface{}) (*model.Employee,error)
	DeleteEmployee(ctx context.Context, cond map[string]interface{}) error
}

type deleteEmployeeBiz struct {
	storage DeleteEmployeeStorage
}

func NewDeleteEmployeeBiz(storage DeleteEmployeeStorage) *deleteEmployeeBiz {
	return &deleteEmployeeBiz{storage: storage}
}

func (biz *deleteEmployeeBiz) DeleteEmployeeById(ctx context.Context, id string) error {
	data, err := biz.storage.GetEmployee(ctx, map[string]interface{}{"manv": id})
	if err != nil {
		return err
	}

	if data.TrangThai != nil && *data.TrangThai == model.EmployeeStatusDeleted {
		return model.ErrEmployeeDeleted
	}

	if err := biz.storage.DeleteEmployee(ctx, map[string]interface{}{"manv": id}); err != nil {
		return err
	}
	return nil
}