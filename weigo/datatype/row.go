package datatype

type Row map[string]interface{}

func (row *Row) Set(key string, val interface{}) {
	(*row)[key] = val
}

func (row *Row) Get(key string) interface{} {
	return (*row)[key]
}
