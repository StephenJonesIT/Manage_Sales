package model

import (
	"manage_sales/common"
)

var EntityName = "Customer"

type Customer struct {
	common.SQLCustomerModel
	LoaiKH  *CustomerType `json:"LoaiKH" gorm:"column:loai_kh;`
}

func (Customer) TableName() string {
	return "khach_hang"
}

type CustomerCreate struct{
	common.SQLCustomerModel
}

func (CustomerCreate) TableName() string{
	return Customer{}.TableName()
}

type CustomerUpdate struct {
	Ho string 				`json:"Ho,omitempty" gorm:"column:ho;"`
	Ten string 				`json:"Ten,omitempty" gorm:"column:ten;"`
	SDT string				`json:"SDT,omitempty" gorm:"column:sdt;"`
	DiaChi string 			`json:"DiaChi,omitempty" gorm:"column:dia_chi;"`
	LoaiKH  *CustomerType 	`json:"LoaiKH,omitempty" gorm:"column:loai_kh;`
}

func (CustomerUpdate) TableName() string{
	return Customer{}.TableName()
}