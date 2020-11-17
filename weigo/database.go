package weigo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	DB        *sql.DB
	TableName string
	Field     string
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
	return database
}

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/platform?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接失败：", err.Error())
	}
	return db
}

func (database *DataBase) Select() {

}
