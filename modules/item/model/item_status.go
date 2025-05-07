package model

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strings"
)

// -------Khai báo enum như sau:
type ItemStatus int

const (
	ItemStatusDoing ItemStatus = iota
	ItemStatusDone
	ItemStatusDeleted
)

var allitemStatus = [3]string{"Doing", "Done", "Deleted"}

func (item *ItemStatus) String() string {
	return allitemStatus[*item]
}

func parseStr2ItemStatus(s string) (ItemStatus, error) {
	for i := range allitemStatus {
		if allitemStatus[i] == s {
			return ItemStatus(i), nil
		}
	}

	return ItemStatus(0), errors.New("invalid status string")
}

// dùng để scan dữ liệu từ DB -> enum
// vì dữ liệu dưới DB và structure hiện tại đang khác nhau.
func (item *ItemStatus) Scan(value interface{}) error {
	// casting value -> []byte
	bytes, ok := value.([]byte)

	if !ok {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	strValue := string(bytes)

	v, err := parseStr2ItemStatus(strValue)
	if err != nil {
		return fmt.Errorf("fail to scan data from sql: %s", value)
	}

	*item = v
	return nil
}

// Structure -> DB
func (item *ItemStatus) Value() (driver.Value, error) {
	if item == nil {
		return nil, nil
	}

	return item.String(), nil
}

// JSON Encoding
func (item *ItemStatus) MarshalJSON() ([]byte, error) {
	if item == nil {
		return nil, nil
	}
	return []byte(fmt.Sprintf("\"%s\"", item.String())), nil
}

// JSON Decoding
func (item *ItemStatus) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")
	itemValue, err := parseStr2ItemStatus(str)
	if err != nil {
		return err
	}
	*item = itemValue
	return nil
}

// ------Kết thúc khai báo enum
