package dialect

import "reflect"

//数据库方言

var dialectMap = map[string]Dialect{}

type Dialect interface {
	//将Go语言数据类型转化为对应的数据库类型
	DataTypeOf(typ reflect.Value) string
	//某个表是否存在的SQL
	TableExistSQL(tableName string) (string, []interface{})
}

func RegisterDialect(name string, d Dialect) {
	dialectMap[name] = d
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectMap[name]
	return
}
