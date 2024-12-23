package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"
)

type CreateEmployeeStorage interface {
	CreateEmployee(context context.Context, data *model.EmployeeCreate) error
}

type createEmployeeBiz struct{
	storage CreateEmployeeStorage
}

func NewCreateEmployeeBiz(storage CreateEmployeeStorage) *createEmployeeBiz{
	return &createEmployeeBiz{storage: storage}
}

func(biz *createEmployeeBiz) CreateNewEmployee(context context.Context, data *model.EmployeeCreate) error{
	if err := biz.storage.CreateEmployee(context,data); err != nil{
		return common.ErrCannotCreateEntity(model.Entity, err)
	}
	return nil
}