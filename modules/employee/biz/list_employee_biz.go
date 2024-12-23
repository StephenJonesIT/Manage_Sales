package biz

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"
)

type ListEmployeeStorage interface {
	ListEmployee(
		ctx context.Context,
		filter *model.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]model.Employee, error)
}

type listEmployeeBiz struct {
	storage ListEmployeeStorage
}

func NewListEmployeeBiz(storage ListEmployeeStorage) *listEmployeeBiz {
	return &listEmployeeBiz{storage: storage}
}

func (binz *listEmployeeBiz) ListEmployeeById(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
) ([]model.Employee, error) {

	data, err := binz.storage.ListEmployee(ctx, filter, paging)

	if err != nil {
		return nil, err
	}
	return data, nil
}