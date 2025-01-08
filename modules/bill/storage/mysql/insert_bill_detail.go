package mysql

import (
	"context"
	"manage_sales/common"
	"manage_sales/modules/bill/model"
)

func (sql *billRepository) InsertDetailBill(ctx context.Context, data *model.ChiTietHoaDon) error{
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}