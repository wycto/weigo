package weigo

/*MVC的M层，模型类*/
type Model struct {
	tableName string
}

func (receiver *Model) SetTableName(Name string) *Model {
	receiver.tableName = Name
	return receiver
}

func (receiver *Model) Find() (map[string]interface{}, string) {
	return DataBase.Name(receiver.tableName).Find()
}

func (receiver *Model) Select() ([]map[string]interface{}, string) {
	return DataBase.Name(receiver.tableName).Select()
}
