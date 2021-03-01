package datatype

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Interfaces []interface{}

func (i *Interfaces) Value() interface{} {
	if len(*i) > 0 {
		return (*i)[0]
	} else {
		return nil
	}
}

func (i *Interfaces) ToString() string {
	v := i.Value()
	if v != nil {
		return fmt.Sprintf("%v", (*i)[0])
	} else {
		return ""
	}
}

func (i *Interfaces) ToInt() int {
	v, err := strconv.Atoi(i.ToString())
	if err != nil {
		return 0
	}

	return v
}

func (i *Interfaces) ToInt8() int8 {
	v, err := strconv.ParseInt(i.ToString(), 10, 8)
	if err != nil {
		return 0
	}

	return int8(v)
}

func (i *Interfaces) ToInt16() int16 {
	v, err := strconv.ParseInt(i.ToString(), 10, 16)
	if err != nil {
		return 0
	}

	return int16(v)
}

func (i *Interfaces) ToInt32() int32 {
	v, err := strconv.ParseInt(i.ToString(), 10, 32)
	if err != nil {
		return 0
	}

	return int32(v)
}

func (i *Interfaces) ToInt64() int64 {
	v, err := strconv.ParseInt(i.ToString(), 10, 64)
	if err != nil {
		return 0
	}

	return v
}

func (i *Interfaces) SetStruct(value interface{}) error {
	jsonStr := i.ToString()
	if len(jsonStr) > 0 {
		err := json.Unmarshal([]byte(jsonStr), value)
		return err
	}
	return errors.New("数据获取失败")
}
