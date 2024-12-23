package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"
)

type EmployeeStorage interface {
	GetEmployee(ctx context.Context, cond map[string]interface{}) (*model.Employee,error)
}

type getEmployeeBiz struct{
	storage EmployeeStorage
}

func NewGetEmployeeBiz(storage EmployeeStorage) (*getEmployeeBiz){
	return &getEmployeeBiz{storage: storage}
}

func(biz *getEmployeeBiz) GetEmployeeById(contex context.Context, id string) (*model.Employee,error){
	data, err := biz.storage.GetEmployee(contex, map[string]interface{}{"manv":id})
	if err != nil {
		return nil, common.ErrCannotGetEntity(model.Entity, err)
	}

	return data, nil
}