package common


type SQLBonsaiModel struct {
	MaSP    string  `json:"MaSP"  gorm:"column:masp; primaryKey"`
	TenSP   string  `json:"TenSP" gorm:"column:tensp;"`
	SoLuong int     `json:"SoLuong" gorm:"column:so_luong;"`
	DonGia  float64 `json:"DonGia" gorm:"column:don_gia;"`
	LoaiCay int     `json:"LoaiCay" gorm:"column:loai_cay;"`
	DVT     string  `json:"DVT,omitempty" gorm:"column:dvt;"`
}

type SQLSuplierModel struct {
	MaNCC   string `json:"MaNCC"  gorm:"column:mancc; primaryKey"`
	Ho      string `json:"Ho" 	gorm:"column:ho;"`
	Ten     string `json:"Ten" gorm:"column:ten;"`
	DiaChi  string `json:"DiaChi" gorm:"column:dia_chi;"`
	SDT		string `json:"SDT" gorm:"column:sdt;"`
}

type SQLCustomerModel struct{
	MaKH string 			`json:"MaKH" gorm:"column:makh;"`
	Ho string 				`json:"Ho" gorm:"column:ho;"`
	Ten string 				`json:"Ten" gorm:"column:ten;"`
	SDT string				`json:"SDT" gorm:"column:sdt;"`
	DiaChi string 			`json:"DiaChi" gorm:"column:dia_chi;"`
}

type Employee struct{
	MaNV string `json:"MaNV" gorm:"column:manv;"`
	Ho string `json:"Ho" gorm:"column:ho;`
	Ten string `json:"Ten" gorm:"column:ten;"`
	DiaChi string `json:"DiaChi" gorm:"column:dia_chi;"`
	SDT string `json:"SDT" gorm:"column:sdt;"`
}

type LoginRequest struct {
	MaTaiKhoan  string `json:"MaTaiKhoan,omitempty" gorm:"column:ma_tai_khoan"`
	TenDangNhap string `json:"TenDangNhap" gorm:"column:ten_dang_nhap"`
	MatKhau     string `json:"MatKhau" gorm:"column:mat_khau"`
}



type HoaDon struct { 
	MaHD string `gorm:"column:mahd" json:"mahd"` 
	TongTien float64 `gorm:"column:tong_tien" json:"tong_tien"` 
	MaKH string `gorm:"index;column:makh" json:"makh"` 
	MaBaoCao string `gorm:"index;column:ma_bao_cao" json:"ma_bao_cao, omitempty"` 
	MaNV string `gorm:"index;column:manv" json:"manv"` 
}

func(HoaDon) TableName() string{
	return "hoa_don"
}