package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

type ImportSlipStatus int

const (
	Doing ImportSlipStatus = iota
	Done
	Deleted
)

var allGoodsReceivedNote = [3]string{"Doing", "Done", "Deleted"}

func (item *ImportSlipStatus) String() string {
	return allGoodsReceivedNote[*item]
}

func parseStrItemStatus(s string) (ImportSlipStatus, error) {
	for i := range allGoodsReceivedNote {
		if allGoodsReceivedNote[i] == s {
			return ImportSlipStatus(i), nil
		}
	}
	return ImportSlipStatus(0), errors.New("invalid status string")
}

func (item *ImportSlipStatus) Scan(value interface{}) error{
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s",value)
	}

	v, err := parseStrItemStatus(string(bytes))
	
	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s",value)
	}

	*item = v
	return nil
}

func (item *ImportSlipStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}
	return item.String(), nil
}

func (item *ImportSlipStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

func (item *ImportSlipStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStrItemStatus(str)

	if err != nil {
		return err
	}

	*item = itemValue

	return nil
}