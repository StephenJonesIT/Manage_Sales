package model

var EntityDetail = "ChiTietPhieuNhap"

type ChiTietPhieuNhap struct{
	MaSP 		string 	`json:"MaSP" gorm:"column:masp"`
	MaPN 		string 	`json:"MaPN" gorm:"column:mapn"`
	SoLuong 	int 	`json:"SoLuong" gorm:"column:so_luong"`
}

func (ChiTietPhieuNhap) TableName() string{
	return "chi_tiet_pn"
}