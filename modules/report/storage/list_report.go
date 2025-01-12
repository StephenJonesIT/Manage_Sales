package storage

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/report/model"
)

func (sql *sqlStore) ListReport(ctx context.Context, paging *common.Paging,morekey ...string) ([]model.Report, error){
	var result []model.Report

	if err := sql.db.Table(model.Report{}.TableName()).Count(&paging.Total).Error; err != nil{
		return nil, err
	}

	if err := sql.db.Order("ngay_tao desc").Offset((paging.Page-1)*paging.Limit).Limit(paging.Limit).Find(&result).Error; err!=nil{
		return nil, err
	}
	return result, nil
}