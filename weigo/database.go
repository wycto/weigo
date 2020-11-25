package weigo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
)

type dataBase struct {
	initStatus bool
	db         *sql.DB
	tableName  string
	fields     string
	where      string
	group      string
	having     string
	order      string
	limit      string
}

func (database *dataBase) getConnect() {
	dataBaseConfig := Config.DB
	db, err := sql.Open(dataBaseConfig.Type, dataBaseConfig.UserName+":"+dataBaseConfig.Password+"@tcp("+dataBaseConfig.HostName+":"+dataBaseConfig.Port+")/"+dataBaseConfig.Database+"?charset="+dataBaseConfig.Charset)
	if err != nil {
		fmt.Println("mysql connect fail", err.Error())
	} else {
		fmt.Println("mysql connect ok")
		database.initStatus = true
	}

	database.db = db
	database.tableName = ""
	database.fields = "*"
	database.where = ""
	database.group = ""
	database.having = ""
	database.order = ""
}

func (database *dataBase) Table(tableName string) *dataBase {
	database.tableName = tableName
	return database
}

func (database *dataBase) Name(tableName string) *dataBase {
	database.tableName = Config.DB.Prefix + tableName
	return database
}

func (database *dataBase) SetFields(fields string) *dataBase {
	database.fields = fields
	return database
}

func (database *dataBase) Where(where interface{}) *dataBase {
	whereStr := ""

	ValueOf := reflect.ValueOf(where)
	valueType := ValueOf.Kind().String()
	if valueType == "map" {
		MapRange := ValueOf.MapRange()
		for MapRange.Next() {
			Key := MapRange.Key().String()
			Value := MapRange.Value().String()

			if string(Value) == Value {
				Value = "\"" + Value + "\""
			}

			if whereStr == "" {
				whereStr = Key + "=" + Value
			} else {
				whereStr += " AND " + Key + "=" + Value
			}
		}
	} else if valueType == "string" {
		whereStr = ValueOf.String()
	} else {

	}

	if whereStr != "" {
		if database.where == "" {
			database.where = " WHERE " + whereStr
		} else {
			database.where += " AND (" + whereStr + ")"
		}
	}

	return database
}

func (database *dataBase) GetOne() (map[string]interface{}, error) {
	rows, err := database.db.Query("SELECT " + database.fields + " FROM " + database.tableName + database.where + database.group + database.having + database.order + " LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return nil, errColumns
	}

	columnLength := len(columns)
	scanByte := make([]interface{}, columnLength) //临时存储每行数据
	values := make([]interface{}, columnLength)   //临时存储每行数据
	for index, _ := range scanByte {              //为每一列初始化一个指针
		scanByte[index] = &values[index]
	}

	rows.Next()
	err = rows.Scan(scanByte...)
	if err != nil {
		return nil, err
	}

	row := make(map[string]interface{})
	for i, data := range values {
		if data != nil {
			row[columns[i]] = string(data.([]byte)) //取实际类型
		} else {
			row[columns[i]] = "" //取实际类型
		}
	}

	return row, nil
}

func (database *dataBase) GetAll() ([]map[string]interface{}, error) {
	SQL := "SELECT " + database.fields + " FROM " + database.tableName + database.where + database.group + database.having + database.order + database.limit
	rows, err := database.db.Query(SQL)
	if err != nil {
		fmt.Println("SQL:", SQL)
		return nil, err
	} else {
		fmt.Println("SQL:", SQL)
	}
	defer rows.Close()

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return nil, errColumns
	}

	columnLength := len(columns)
	scanByte := make([]interface{}, columnLength)
	values := make([]interface{}, columnLength)
	for index, _ := range scanByte {
		scanByte[index] = &values[index]
	}

	var list []map[string]interface{}
	for rows.Next() {
		err := rows.Scan(scanByte...)
		if err != nil {
			return nil, err
		}

		item := make(map[string]interface{})
		for i, data := range values {
			if data != nil {
				item[columns[i]] = string(data.([]byte))
			} else {
				item[columns[i]] = ""
			}
		}
		list = append(list, item)
	}
	return list, nil
}
