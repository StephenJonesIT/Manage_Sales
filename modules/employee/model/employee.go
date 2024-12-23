package model

import (
	"errors"
	"manage_sales/common"
)
var Entity = "NhanVien"

var (
	ErrEmployeeDeleted = errors.New("Employee is deleted")
)

type Employee struct {
	common.Employee
	TrangThai *EmployeeStatus `json:"TrangThai" gorm:"column:trang_thai"`
}

func (Employee) TableName() string {
	return "nhan_vien"
}

type EmployeeCreate struct {
	common.Employee
}

func (EmployeeCreate) TableName() string{
	return Employee{}.TableName()
}

type EmployeeUpdate struct {
	Ho        string 			`json:"Ho,omitempty" gorm:"column:ho;`
	Ten       string 			`json:"Ten,omitempty" gorm:"column:ten;"`
	DiaChi    string 			`json:"DiaChi,omitempty" gorm:"column:dia_chi;"`
	SDT       string 			`json:"SDT,omitempty" gorm:"column:sdt;"`
	TrangThai *EmployeeStatus 	`json:"TrangThai,omitempty" gorm:"column:trang_thai"`
}

func (EmployeeUpdate) TableName() string{
	return Employee{}.TableName()
}