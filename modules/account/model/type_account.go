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

func (item *AccountType) String() string {
	return allAccountType[*item]
}

func parseStrAccountType(s string) (AccountType, error) {
	for i := range allAccountType {
		if allAccountType[i]==s {
			return AccountType(i), nil
		}
	}
	return AccountType(0), errors.New("Invalid type string")
}

func (item *AccountType) Scan(value interface{}) error{
	bytes, ok := value.([]byte) 
	if !ok {
		return fmt.Errorf("fail to scan data from sql:  %s", ok)
	}

	data, err := parseStrAccountType(string(bytes))

	if err != nil {
		return fmt.Errorf("fail to scan data from sql:  %s", err)
	}

	*item = data

	return nil
}


func (item *AccountType) MarshalJson() ([]byte, error){
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}


func (item *AccountType) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *AccountType) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStrAccountType(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}