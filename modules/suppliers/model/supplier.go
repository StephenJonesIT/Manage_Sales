package model

import (
	"errors"
	"manage_sales/common"
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
}

func(SupplierItem) TableName() string{
	return "nha_cung_cap"
}

type SupplierItemCreate struct {
	common.SQLSuplierModel
	LoaiNCC 	int    		`json:"LoaiNCC" gorm:"column:loai_ncc;"`
}

func (SupplierItemCreate) TableName() string{
	return SupplierItem{}.TableName();
}

type SupplierItemUpdate struct {
	Ho			string  			`json:"Ho,omitempty" gorm:"column:ho;"`
	Ten   	  	string     			`json:"Ten,omitempty" gorm:"column:ten;"`
	DiaChi    	string 				`json:"DiaChi,omitempty" gorm:"column:dia_chi;"`
	LoaiNCC   	int     			`json:"LoaiNCC,omitempty" gorm:"column:loai_ncc;"`
	SDT 		string 				`json:"SDT,omitempty" gorm:"column:sdt;"`
}

func (SupplierItemUpdate) TableName() string{
	return SupplierItem{}.TableName()
}