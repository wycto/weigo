package datatype

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Row map[string]interface{}

func (row *Row) Set(key string, val interface{}) {
	(*row)[key] = val
}

func (row *Row) Get(key string) interface{} {
	return (*row)[key]
}

func (row *Row) Has(key string) bool {
	_, ok := (*row)[key]
	if ok {
		return true
	}
	return false
}

func (row *Row) GetString(key string) string {
	if row.Has(key) {
		return fmt.Sprintf("%v", row.Get(key))
	} else {
		return ""
	}
}

func (row *Row) GetInt(key string) int {
	v, err := strconv.Atoi(row.GetString(key))
	if err != nil {
		return 0
	}

	return v
}

func (row *Row) GetInt8(key string) int8 {
	v, err := strconv.ParseInt(row.GetString(key), 10, 8)
	if err != nil {
		return 0
	}

	return int8(v)
}

func (row *Row) GetInt16(key string) int16 {
	v, err := strconv.ParseInt(row.GetString(key), 10, 16)
	if err != nil {
		return 0
	}

	return int16(v)
}

func (row *Row) GetInt32(key string) int32 {
	v, err := strconv.ParseInt(row.GetString(key), 10, 32)
	if err != nil {
		return 0
	}

	return int32(v)
}

func (row *Row) GetInt64(key string) int64 {
	v, err := strconv.ParseInt(row.GetString(key), 10, 64)
	if err != nil {
		return 0
	}

	return v
}

func (row *Row) SetStruct(key string, value interface{}) error {
	jsonStr := row.GetString(key)
	if len(jsonStr) > 0 {
		err := json.Unmarshal([]byte(jsonStr), value)
		return err
	}
	return errors.New("数据获取失败")
}
