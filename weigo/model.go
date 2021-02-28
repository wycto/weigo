package weigo

import "wycto/weigo/datatype"

/*MVC的M层，模型类*/
type Model struct {
	tableName string
}

func (model *Model) SetTableName(Name string) *Model {
	model.tableName = Name
	return model
}

func (model *Model) Where(where *datatype.Row) *Model {
	DB.Where(where)
	return model
}

func (model *Model) Find() (row *datatype.Row, err error) {
	return DB.Name(model.tableName).Find()
}

func (model *Model) Select() (rows *datatype.Rows, err error) {
	return DB.Name(model.tableName).Select()
}

func (model *Model) Insert(data *datatype.Row) (id string, err error) {
	return DB.Insert(data)
}
