package mysql

import (
	"context"
	"manage_sales/modules/bill/model"
)

func (sql *billRepository) DeleteBill(ctx context.Context, cond map[string]interface{}) error {
	deleteBill := model.Deleted

	if err := sql.db.Table(model.HoaDon{}.TableName()).Where(cond).Updates(map[string] interface{}{
		"trang_thai":deleteBill.String(),
	}).Error; err != nil{
		return err
	}

	return nil
}