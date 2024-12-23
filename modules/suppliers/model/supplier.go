package model

import (
	"errors"
	"manage_sales/common"
	"time"
)
const (
	EntityName = "Bonsai"
)

var (
	ErrTitleIsBlank = errors.New("ten cannot be blank")
	ErrItemDeleted  = errors.New("supplier is deleted")
)
type SupplierItem struct {
	common.SQLSuplierModel
	LoaiNCC 	int    				`json:"LoaiNCC" gorm:"column:loai_ncc;"`
	TrangThai 	*SupplierStatus  	`json:"TrangThai" gorm:"trang_thai"`
	NgayTao 	*time.Time			`json:"NgayTao" gorm:"ngay_tao"`
}

func(SupplierItem) TableName() string{
	return "nha_cung_cap"
}

type SupplierItemCreate struct {
	common.SQLSuplierModel
}

func (SupplierItemCreate) TableName() string{
	return SupplierItem{}.TableName();
}

type SupplierItemUpdate struct {
	Ho			string  			`json:"Ho,omitempty" gorm:"column:ho;"`
	Ten   	  	string     			`json:"Ten,,omitempty" gorm:"column:ten;"`
	DiaChi    	string 				`json:"DiaChi,,omitempty" gorm:"column:dia_chi;"`
	LoaiNCC   	int     			`json:"LoaiNCC,,omitempty" gorm:"column:loai_ncc;"`
}

func (SupplierItemUpdate) TableName() string{
	return SupplierItem{}.TableName()
}