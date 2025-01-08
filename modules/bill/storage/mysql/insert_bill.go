package mysql

import (
	"context"
	"manage_sales/common"
)

func (sql *billRepository) InsertBill(ctx context.Context, data *common.HoaDon) error{
	if err := sql.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}