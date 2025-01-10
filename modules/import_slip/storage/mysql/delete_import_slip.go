package mysql

import (
	"context"
	"manage_sales/modules/import_slip/model"
)

func (sql *sqlStore) DeleteImportSlip(ctx context.Context, cond map[string]interface{}) error {
	deleteImportSlip := model.Deleted

	if err := sql.db.Table(model.PhieuNhap{}.TableName()).Where(cond).Updates(map[string] interface{}{
		"trang_thai":deleteImportSlip.String(),
	}).Error; err != nil{
		return err
	}

	return nil
}