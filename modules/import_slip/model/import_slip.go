package model

import "time"
var EntitySlip = "PhieuNhap"
type PhieuNhap struct {
	MaPN      string             `json:"mapn" gorm:"column:mapn"`
	TongTien  int                `json:"TongTien" gorm:"column:tong_tien"`
	NgayLap   *time.Time         `json:"NgayLap,omitempty" gorm:"column:ngay_lap_pn"`
	NgaySua   *time.Time         `json:"NgaySua,omitempty" gorm:"column:ngay_chinh_sua"`
	MaNCC     string             `json:"MaNCC" gorm:"column:mancc"`
	MaBaoCao  string             `json:"MaBaoCao,omitempty" gorm:"column:"ma_bao_cao"`
	MaNV      string             `json:"MaNV,omitempty" gorm:"column:manv"`
	TrangThai *ImportSlipStatus  `json:"TrangThai,omitempty" gorm:"column:trang_thai"`
	ChiTiet   []ChiTietPhieuNhap `json:"ChiTiet" gorm:"foreignKey:MaPN;references:MaPN"`
}

func (PhieuNhap) TableName() string {
	return "phieu_nhap"
}

type PhieuNhapCreate struct{
	MaPN      string             `json:"mapn" gorm:"column:mapn"`
	TongTien  int                `json:"TongTien" gorm:"column:tong_tien"`
	MaNCC     string             `json:"MaNCC" gorm:"column:mancc"`
	MaBaoCao  string             `json:"MaBaoCao,omitempty" gorm:"column:"ma_bao_cao"`
	MaNV      string             `json:"MaNV,omitempty" gorm:"column:manv"`
	ChiTiet   []ChiTietPhieuNhap `json:"ChiTiet" gorm:"foreignKey:MaPN;references:MaPN"`
}

func(PhieuNhapCreate) TableName() string{
	return PhieuNhap{}.TableName()
}

type PhieuNhapUpdate struct{
	MaPN      string             `json:"mapn" gorm:"column:mapn"`
	TongTien  int                `json:"TongTien" gorm:"column:tong_tien"`
	MaNCC     string             `json:"MaNCC" gorm:"column:mancc"`
	MaBaoCao  string             `json:"MaBaoCao,omitempty" gorm:"column:"ma_bao_cao"`
	MaNV      string             `json:"MaNV,omitempty" gorm:"column:manv"`
	TrangThai *ImportSlipStatus  `json:"TrangThai,omitempty" gorm:"column:trang_thai"`
	ChiTiet   []ChiTietPhieuNhap `json:"ChiTiet" gorm:"foreignKey:MaPN;references:MaPN"`
}