package utils

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Date struct {
	time.Time
}

// MarshalJSON 序列化为JSON
func (t Date) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02"))
	return []byte(stamp), nil
}

// UnmarshalJSON 反序列化为JSON
func (t *Date) UnmarshalJSON(data []byte) error {
	var err error
	t.Time, err = time.Parse("2006-01-02", string(data)[1:11])
	return err
}

// String 重写String方法
func (t *Date) String() string {
	data, _ := json.Marshal(t)
	return string(data)
}

// FieldType 数据类型
func (t *Date) FieldType() int {
	return orm.TypeDateTimeField

}

// SetRaw 读取数据库值
func (t *Date) SetRaw(value interface{}) error {
	switch value.(type) {
	case time.Time:
		t.Time = value.(time.Time)
	}
	return nil
}

// RawValue 写入数据库
func (t *Date) RawValue() interface{} {
	str := t.Format("2006-01-02")
	if str == "0001-01-01" {
		return nil
	}
	return str
}
