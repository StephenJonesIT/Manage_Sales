package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/employee/model"
)

func (s *sqlStore) ListEmployee(
	ctx context.Context,
	filter *model.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]model.Employee, error) {

	var result []model.Employee

	db := s.db.Where("trang_thai <> ?", "Deleted")

	if f := filter; f != nil {
		if v := f.TrangThai; v != "" {
			db = db.Where("trang_thai = ?", v)
		}
	}

	if err := db.Table(model.Employee{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if err := db.Order("manv desc").
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}