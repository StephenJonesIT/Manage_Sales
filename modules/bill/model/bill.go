package model

import "time"

var Entity = "HoaDon"

type HoaDon struct {
	MaHD      string          `gorm:"column:mahd" json:"MaHD"`
	TongTien  float64         `gorm:"column:tong_tien" json:"TongTien"`
	MaKH      string          `gorm:"index;column:makh" json:"MaKH"`
	MaBaoCao  string      	  `gorm:"index;column:ma_bao_cao" json:"MaBaoCao,omitempty"`
	MaNV      string          `gorm:"index;column:manv" json:"MaNVMaNV,omitempty"`
	NgaySua   *time.Time      `json:"NgaySua,omitempty" gorm:"column:ngay_chinh_sua"`
	NgayLap   *time.Time      `json:"NgayLap,omitempty" gorm:"column:ngay_lap_hdhd"`
	TrangThai *BillStatus     `json:"TrangThai" gorm:"column:trang_thai"`
	ChiTiet   []ChiTietHoaDon `gorm:"foreignKey:MaHD;references:MaHD" json:"ChiTiet"`
}

func (HoaDon) TableName() string {
	return "hoa_don"
}

type CreateHoaDon struct {
	MaHD      string          `gorm:"column:mahd" json:"MaHD"`
	TongTien  float64         `gorm:"column:tong_tien" json:"TongTien"`
	MaKH      string          `gorm:"column:makh" json:"MaKH"`
	MaBaoCao  string      	  `gorm:"column:ma_bao_cao" json:"MaBaoCao,omitempty"`
	MaNV      string          `gorm:"column:manv" json:"manv,omitempty"`
	ChiTiet   []ChiTietHoaDon `gorm:"foreignKey:MaHD;references:MaHD" json:"ChiTiet"`
}

func (CreateHoaDon) TableName() string {
	return HoaDon{}.TableName()
}

type UpdateHoaDon struct {
	MaHD      string          `gorm:"column:mahd" json:"MaHD"`
	TongTien  int         	  `gorm:"column:tong_tien" json:"TongTien"`
	MaKH      string          `gorm:"column:makh" json:"MaKH"`
	MaBaoCao  string      	  `gorm:"column:ma_bao_cao" json:"MaBaoCao,omitempty"`
	MaNV      string          `gorm:"column:manv" json:"manv,omitempty"`
	NgaySua   *time.Time      `json:"NgaySua,omitempty" gorm:"column:ngay_chinh_sua"`
	TrangThai *BillStatus     `json:"TrangThai" gorm:"column:trang_thai"`
	ChiTiet   []ChiTietHoaDon `gorm:"foreignKey:MaHD;references:MaHD" json:"ChiTiet"`
}

func(UpdateHoaDon) TableName() string{
	return HoaDon{}.TableName()
}