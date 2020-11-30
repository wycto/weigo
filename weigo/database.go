package weigo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"regexp"
	"strconv"
	"strings"
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
		fmt.Println("mysql open ok")
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
	database.tableName = "`" + tableName + "`"
	return database
}

func (database *dataBase) Name(tableName string) *dataBase {
	database.tableName = "`" + Config.DB.Prefix + tableName + "`"
	return database
}

func (database *dataBase) SetFields(fieldsStr string) *dataBase {
	fieldsStr = strings.Replace(fieldsStr, "`", "", -1)
	fieldsStr = strings.Replace(fieldsStr, ",", "`,`", -1)
	fieldsStr = "`" + fieldsStr + "`"
	database.fields = fieldsStr
	return database
}

func (database *dataBase) Where(where interface{}) *dataBase {
	whereStr := ""
	ValueOf := reflect.ValueOf(where)
	valueType := ValueOf.Kind().String()
	if valueType == "map" {
		MapRange := ValueOf.MapRange()
		for MapRange.Next() {

			field, reg := database.getReflectKey(MapRange.Key().String())
			field = strings.Replace(field, "`", "", -1)
			field = "`" + field + "`"

			Value := database.getReflectValue(MapRange.Value().Interface())

			if whereStr == "" {
				whereStr = field + reg + Value
			} else {
				whereStr += " AND " + field + reg + Value
			}
		}
	} else if valueType == "string" {
		whereStr = ValueOf.String()
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

func (database *dataBase) Group(groupStr string) *dataBase {
	database.group = " GROUP " + groupStr
	return database
}

func (database *dataBase) Having(havingStr string) *dataBase {
	database.having = " HAVING " + havingStr
	return database
}

func (database *dataBase) Order(orderStr string) *dataBase {
	database.order = " ORDER " + orderStr
	return database
}

func (database *dataBase) Limit(limitStr string) *dataBase {
	database.limit = " LIMIT " + limitStr
	return database
}

func (database *dataBase) Page(page int, count int) *dataBase {
	begin := (page - 1) * count
	database.limit = " LIMIT " + strconv.Itoa(begin) + "," + strconv.Itoa(count)
	return database
}

func (database *dataBase) GetOne() (map[string]interface{}, string) {
	rows, err := database.db.Query("SELECT " + database.fields + " FROM " + database.tableName + database.where + database.group + database.having + database.order + " LIMIT 1")
	database.resetSelectInfo()

	if err != nil {
		return nil, database.getErrorString(err.Error())
	}
	defer rows.Close()

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return nil, errColumns.Error()
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
		return nil, err.Error()
	}

	row := make(map[string]interface{})
	for i, data := range values {
		if data != nil {
			row[columns[i]] = string(data.([]byte)) //取实际类型
		} else {
			row[columns[i]] = "" //取实际类型
		}
	}

	return row, ""
}

func (database *dataBase) GetAll() ([]map[string]interface{}, string) {
	SQL := "SELECT " + database.fields + " FROM " + database.tableName + database.where + database.group + database.having + database.order + database.limit
	database.resetSelectInfo()

	rows, err := database.db.Query(SQL)
	if err != nil {
		fmt.Println("SQL:", SQL)
		return nil, database.getErrorString(err.Error())
	} else {
		if Config.Log.SqlInfo == "console" {
			fmt.Println(Log.FormatLogString(SQL, "Info", "SQL"))
		}
	}
	defer rows.Close()

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return nil, errColumns.Error()
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
			return nil, err.Error()
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
	return list, ""
}

func (database *dataBase) resetSelectInfo() {
	database.where = ""
}

func (database *dataBase) getErrorString(ErrorStr string) string {
	length := len(ErrorStr)
	if length > 11 {
		if ErrorStr[:10] == "Error 1045" {
			return "数据库连接失败，请检查连接账号密码：" + ErrorStr
		} else if ErrorStr[:10] == "Error 1049" {
			return "数据库不存在：" + ErrorStr
		} else if ErrorStr[:8] == "dial tcp" {
			return "数据库连接失败，请检查网络：" + ErrorStr
		}

	}
	return ErrorStr
}

func (database *dataBase) getReflectKey(fieldReg string) (string, string) {
	reg, err := regexp.Compile(`\|(.*)`)
	if err != nil {
		fmt.Println("regexp err:", err.Error())
		return fieldReg, "="
	}

	result := reg.FindAllString(fieldReg, -1)
	if len(result) == 0 {
		return fieldReg, "="
	}

	fieldRegIndexArr := reg.FindAllStringIndex(fieldReg, -1)
	position := fieldRegIndexArr[0][0]
	field := fieldReg[:position]
	regexpStr := fieldReg[position+1:]
	regexpStr = strings.Trim(regexpStr, " ")
	if regexpStr == "" {
		regexpStr = "="
	}
	return field, regexpStr
}

func (database *dataBase) getReflectValue(InterfaceValue interface{}) string {
	Value := "\"\""
	ValueInterfaceValueOf := reflect.ValueOf(InterfaceValue)
	valueType := ValueInterfaceValueOf.Kind()
	fmt.Println("Type:", valueType)
	fmt.Println("Value:", ValueInterfaceValueOf)
	fmt.Println("ValueString:", ValueInterfaceValueOf.String())
	switch valueType {
	case reflect.String:
		Value = "\"" + strings.Trim(ValueInterfaceValueOf.String(), " ") + "\""
	case reflect.Float64:
		Value = strconv.FormatFloat(ValueInterfaceValueOf.Float(), 'E', -1, 64)
	case reflect.Float32:
		Value = strconv.FormatFloat(ValueInterfaceValueOf.Float(), 'E', -1, 32)
	case reflect.Int:
		Value = strconv.Itoa(int(ValueInterfaceValueOf.Int()))
	case reflect.Int64:
		Value = strconv.FormatInt(ValueInterfaceValueOf.Int(), 10)
	case reflect.Bool:
		Value = strconv.FormatBool(ValueInterfaceValueOf.Bool())
	default:
		fmt.Println(Log.FormatLogString("SELECT WHERE Parse Error; Value:"+ValueInterfaceValueOf.String(), "Error", "SQL"))
	}

	return Value
}
