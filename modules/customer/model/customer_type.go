package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type CustomerType int

const (
	CustomerNormail CustomerType = iota
	CustomerVip
)

var allCustomerType = [2]string{"Normal", "VIP"}

func (item *CustomerType) String() string {
	return allCustomerType[*item]
}

func parseStr2ItemStatus(s string) (CustomerType, error) {
	for i := range allCustomerType {
		if allCustomerType[i] == s {
			return CustomerType(i), nil
		}
	}
	return CustomerType(0), errors.New("invalid status string")
}

// Phương thức này dùng để chuyển đổi kết quả truy vấn SQL thành các giá trị ItemStatus. Nó đọc giá trị dưới dạng []byte,
// chuyển đổi thành chuỗi và sau đó sử dụng hàm parseStr2ItemStatus để chuyển chuỗi thành ItemStatus.
func (item *CustomerType) Scan(value interface{}) error {
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	v, err := parseStr2ItemStatus(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*item = v

	return nil
}

// Phương thức này chuyển một giá trị ItemStatus thành đại diện JSON của nó.
// Nó trả về đại diện dạng chuỗi của giá trị ItemStatus dưới dạng một mảng byte JSON mã hóa.
func (item *CustomerType) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// Mục đích: Hàm này chuyển giá trị của một ItemStatus thành driver.Value (dạng giá trị mà database driver hiểu được).
// Logic:
// Nếu item là nil, trả về nil.
// Ngược lại, trả về chuỗi đại diện của item (bằng cách gọi item.String()).
func (item *CustomerType) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

// Mục đích: Hàm này chuyển đổi dữ liệu JSON thành giá trị ItemStatus.
// Logic:
// Chuyển dữ liệu JSON (dưới dạng []byte) thành chuỗi và loại bỏ các ký tự dấu nháy kép (").
// Gọi hàm parseStr2ItemStatus để chuyển chuỗi đó thành ItemStatus.
// Nếu xảy ra lỗi trong quá trình chuyển đổi, trả về lỗi đó.
// Ngược lại, gán giá trị ItemStatus đã chuyển đổi cho item.
func (item *CustomerType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}