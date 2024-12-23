package model

import (
	"errors"
	"manage_sales/common"
)

const (
	EntityName = "Bonsai"
)

var (
	ErrTitleIsBlank = errors.New("title cannot be blank")
	ErrItemDeleted  = errors.New("item is deleted")
)

type BonsaiItem struct {
	common.SQLBonsaiModel
	TrangThai *BonsaiStatus `json:"TrangThai,omitempty" gorm:"column:trang_thai;"`
}

func (BonsaiItem) TableName() string {
	return "san_pham"
}

type BonsaiItemCreate struct {
	common.SQLBonsaiModel
}

func (BonsaiItemCreate) TableName() string {
	return BonsaiItem{}.TableName()
}

type BonsaiItemUpdate struct {
	TenSP     string  		`json:"TenSP,omitempty" gorm:"column:tensp;"`
	SoLuong   int     		`json:"SoLuong,omitempty" gorm:"column:so_luong;"`
	DonGia    float64 		`json:"DonGia,omitempty" gorm:"column:don_gia;"`
	LoaiCay   int     		`json:"LoaiCay,omitempty" gorm:"column:loai_cay;"`
	DVT       string  		`json:"DVT,omitempty" gorm:"column:dvt;"`
	TrangThai *BonsaiStatus `json:"TrangThai,omitempty" gorm:"column:trang_thai;"`
}

func (BonsaiItemUpdate) TableName() string {
	return BonsaiItem{}.TableName()
}