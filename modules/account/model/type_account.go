package model

import (
    "database/sql/driver"
    "errors"
    "fmt"
    "strings"
)

type AccountType int

const (
    AccountUser AccountType = iota
    AccountAdmin
)

var allAccountType = [2]string{"user", "admin"}

// Phương thức String để chuyển AccountType thành chuỗi
func (item AccountType) String() string {
    if int(item) < len(allAccountType) {
        return allAccountType[item]
    }
    return "unknown"
}

// Chuyển chuỗi thành AccountType
func parseStrAccountType(s string) (AccountType, error) {
    for i, v := range allAccountType {
        if v == s {
            return AccountType(i), nil
        }
    }
    return AccountType(0), errors.New("Invalid type string")
}

// Scan để chuyển dữ liệu từ cơ sở dữ liệu thành AccountType
func (item *AccountType) Scan(value interface{}) error {
    bytes, ok := value.([]byte)
    if !ok {
        return fmt.Errorf("fail to scan data from sql: %t", ok)
    }

    data, err := parseStrAccountType(string(bytes))
    if err != nil {
        return fmt.Errorf("fail to scan data from sql: %s", err)
    }

    *item = data

    return nil
}

// Chuyển đổi AccountType thành giá trị JSON
func (item AccountType) MarshalJSON() ([]byte, error) {
    return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// Chuyển AccountType thành giá trị có thể lưu trữ trong cơ sở dữ liệu
func (item AccountType) Value() (driver.Value, error) {
    return item.String(), nil
}

// Chuyển đổi JSON thành AccountType
func (item *AccountType) UnmarshalJSON(data []byte) error {
    str := strings.ReplaceAll(string(data), "\"", "")
    itemValue, err := parseStrAccountType(str)

    if err != nil {
        return err
    }

    *item = itemValue

    return nil
}
