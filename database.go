package weigo

/*
数据库类
*/
import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wycto/weigo/datatype"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type dataBase struct {
	initStatus bool    //是否初始化类
	db         *sql.DB //db
	tableName  string  //操作的表名称，全名称
	fields     string  //查询的字段
	where      string  //条件，包括查询、更新、删除
	group      string  //分组
	having     string  //
	order      string  //排序
	limit      string  //限制条数
}

//连接数据库
func (database *dataBase) getConnect() {
	dataBaseConfig := Config.DB
	db, err := sql.Open(dataBaseConfig.Type, dataBaseConfig.UserName+":"+dataBaseConfig.Password+"@tcp("+dataBaseConfig.HostName+":"+dataBaseConfig.Port+")/"+dataBaseConfig.Database+"?charset="+dataBaseConfig.Charset)
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库打开成功")
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

//级联操作-设置操作的表名称，全表名称，包含前缀
func (database *dataBase) Table(tableName string) *dataBase {
	database.tableName = "`" + tableName + "`"
	return database
}

//级联操作-设置操作的表名称，不带前缀，使用配置里面的前缀
func (database *dataBase) Name(tableName string) *dataBase {
	database.tableName = "`" + Config.DB.Prefix + tableName + "`"
	return database
}

//级联操作-设置要查询的字段
func (database *dataBase) Fields(fieldsStr string) *dataBase {
	fieldsStr = strings.Replace(fieldsStr, "`", "", -1)
	fieldsStr = strings.Replace(fieldsStr, ",", "`,`", -1)
	fieldsStr = "`" + fieldsStr + "`"
	database.fields = fieldsStr
	return database
}

//级联操作-操作的条件，包含查询、修改、删除
func (database *dataBase) Where(where *datatype.Row) *dataBase {
	whereStr := ""

	for key, val := range *where {
		if whereStr == "" {
			whereStr = key + "='" + fmt.Sprintf("%v", val) + "'"
		} else {
			whereStr += " AND " + key + "='" + fmt.Sprintf("%v", val) + "'"
		}
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

//级联操作-分组
func (database *dataBase) Group(groupStr string) *dataBase {
	database.group = " GROUP " + groupStr
	return database
}

//级联操作-聚合判断
func (database *dataBase) Having(havingStr string) *dataBase {
	database.having = " HAVING " + havingStr
	return database
}

//级联操作-排序
func (database *dataBase) Order(orderStr string) *dataBase {
	database.order = " ORDER " + orderStr
	return database
}

//级联操作-限制条数
func (database *dataBase) Limit(limitStr string) *dataBase {
	database.limit = " LIMIT " + limitStr
	return database
}

//级联操作-分页设置
func (database *dataBase) Page(page int, count int) *dataBase {
	begin := (page - 1) * count
	database.limit = " LIMIT " + strconv.Itoa(begin) + "," + strconv.Itoa(count)
	return database
}

//查询一条数据
func (database *dataBase) Find() (row *datatype.Row, err error) {

	row = &datatype.Row{}

	SQL := "SELECT " + database.fields + " FROM " + database.tableName + database.where + database.group + database.having + database.order + " LIMIT 1"
	rows, err := database.db.Query(SQL)
	if Config.Log.SqlInfo == "console" {
		fmt.Println(Log.FormatLogString(SQL, "Info", "SQL"))
	}
	database.resetSQL()

	if err != nil {
		return row, errors.New(database.getErrorString(err.Error()))
	}
	defer rows.Close()

	columns, errColumns := rows.Columns()
	if errColumns != nil {
		return row, errColumns
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
		return row, err
	}

	for i, data := range values {
		if data != nil {
			(*row)[columns[i]] = string(data.([]byte)) //取实际类型
		} else {
			(*row)[columns[i]] = "" //取实际类型
		}
	}

	return row, err
}

//查询多条数据
func (database *dataBase) Select() (rows *datatype.Rows, err error) {
	rows = &datatype.Rows{}

	SQL := "SELECT " + database.fields + " FROM " + database.tableName + database.where + database.group + database.having + database.order + database.limit
	database.resetSQL()

	dbRows, err := database.db.Query(SQL)
	if err != nil {
		fmt.Println("SQL:", SQL)
		return rows, errors.New(database.getErrorString(err.Error()))
	} else {
		if Config.Log.SqlInfo == "console" {
			fmt.Println(Log.FormatLogString(SQL, "Info", "SQL"))
		}
	}
	defer dbRows.Close()

	columns, errColumns := dbRows.Columns()
	if errColumns != nil {
		return rows, errColumns
	}

	columnLength := len(columns)
	scanByte := make([]interface{}, columnLength)
	values := make([]interface{}, columnLength)
	for index, _ := range scanByte {
		scanByte[index] = &values[index]
	}

	for dbRows.Next() {
		err = dbRows.Scan(scanByte...)
		if err != nil {
			return rows, err
		}

		item := &datatype.Row{}
		for i, data := range values {
			if data != nil {
				(*item)[columns[i]] = string(data.([]byte))
			} else {
				(*item)[columns[i]] = ""
			}
		}
		tmpRows := append(*rows, item)
		rows = &tmpRows
	}
	return rows, nil
}

//插入一条数据
func (database *dataBase) Insert(data *datatype.Row) (id string, err error) {
	insertData := database.getInsertValue(data)
	if insertData == "" {
		return id, errors.New("没有要插入的数据")
	}

	SQL := "INSERT INTO " + database.tableName + insertData
	database.resetSQL()
	result, err := database.db.Exec(SQL)
	if err != nil {
		return id, err
	}

	num, err := result.RowsAffected()
	if err != nil {
		return strconv.Itoa(int(num)), err
	}
	return id, err
}

//更新所有-不使用条件
func (database *dataBase) UpdateAll(data map[string]interface{}) (int64, string) {
	SQL := "UPDATE " + database.tableName + " SET " + database.getUpdateValue(data)
	database.resetSQL()
	result, err := database.db.Exec(SQL)
	if err != nil {
		return 0, err.Error()
	}

	num, err := result.RowsAffected()
	if err != nil {
		return 0, err.Error()
	}
	return num, ""
}

//根据条件更新
func (database *dataBase) Update(data map[string]interface{}) (int64, string) {
	if database.where != "" {
		SQL := "UPDATE " + database.tableName + " SET " + database.getUpdateValue(data) + database.where
		database.resetSQL()
		result, err := database.db.Exec(SQL)
		if err != nil {
			return 0, err.Error()
		}

		num, err := result.RowsAffected()
		if err != nil {
			return 0, err.Error()
		}
		return num, ""
	}

	return 0, "where empty"
}

//删除所有-不使用条件
func (database *dataBase) DeleteAll() (int64, string) {
	SQL := "DELETE FROM " + database.tableName
	database.resetSQL()
	result, err := database.db.Exec(SQL)
	if err != nil {
		return 0, err.Error()
	}

	num, err := result.RowsAffected()
	if err != nil {
		return 0, err.Error()
	}
	return num, ""
}

//根据条件删除
func (database *dataBase) Delete() (int64, string) {
	if database.where != "" {
		SQL := "DELETE FROM " + database.tableName + database.where
		database.resetSQL()
		result, err := database.db.Exec(SQL)
		if err != nil {
			return 0, err.Error()
		}

		num, err := result.RowsAffected()
		if err != nil {
			return 0, err.Error()
		}
		return num, ""
	}

	return 0, "where empty"
}

func (database *dataBase) resetSQL() {
	database.where = ""
}

func (database *dataBase) getErrorString(ErrorStr string) string {
	length := len(ErrorStr)
	if length > 11 {
		if ErrorStr[:10] == "Error 1045" {
			return "数据库连接失败，请检查连接账号密码：" + ErrorStr
		} else if ErrorStr[:10] == "Error 1049" {
			return "数据库不存在：" + ErrorStr
		} else if ErrorStr[:10] == "Error 1146" {
			return "数据表不存在：" + ErrorStr
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

func (database *dataBase) getUpdateValue(updateData map[string]interface{}) string {
	updateValueStr := ""

	for field, data := range updateData {
		value := database.getReflectValue(data)
		updateValueStr += "`" + field + "`=" + value
	}
	return updateValueStr
}

func (database *dataBase) getInsertValue(updateData *datatype.Row) string {
	fields := ""
	values := ""

	for field, data := range *updateData {
		value := database.getReflectValue(data)
		fields += ",`" + field + "`"
		values += "," + value
	}

	if fields == "" {
		return ""
	}

	fields = fields[1:]
	values = values[1:]

	return "(" + fields + ") VALUES (" + values + ")"
}
