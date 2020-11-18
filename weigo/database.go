package weigo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	DB        *sql.DB
	TableName string
	Fields    string
	Where     string
}

func init() {

}

func GetDataBase() *DataBase {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/platform?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败：", err.Error())
	}
	database := &DataBase{}
	database.DB = db
	database.Fields = "*"
	return database
}

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/platform?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败：", err.Error())
	}
	return db
}

func (database *DataBase) Table(tableName string) *DataBase {
	database.TableName = tableName
	return database
}

func (database *DataBase) SetFields(fields string) *DataBase {
	database.Fields = fields
	return database
}

func (database *DataBase) GetRows() ([]map[string]interface{}, error) {
	rows, err := database.DB.Query("SELECT " + database.Fields + " FROM " + database.TableName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return nil, errColumns
	}

	columnTypes, errColumnTypes := rows.ColumnTypes()
	if errColumnTypes != nil {
		return nil, errColumnTypes
	}

	columnLength := len(columns)
	scanByte := make([]interface{}, columnLength) //临时存储每行数据
	values := make([]interface{}, columnLength)   //临时存储每行数据
	for index, _ := range scanByte {              //为每一列初始化一个指针
		scanByte[index] = &values[index]
	}

	var list []map[string]interface{} //返回的切片
	for rows.Next() {
		err := rows.Scan(scanByte...)
		if err != nil {
			return nil, err
		}

		item := make(map[string]interface{})
		for i, data := range values {
			fmt.Println("columnTypes_Name", columnTypes[i].Name())
			fmt.Println("columnTypes_DatabaseTypeName", columnTypes[i].DatabaseTypeName())
			if data != nil {
				item[columns[i]] = string(data.([]byte)) //取实际类型
			} else {
				item[columns[i]] = "" //取实际类型
			}
		}
		list = append(list, item)
	}
	rows.Close()
	return list, nil
}
