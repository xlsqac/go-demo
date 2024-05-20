package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

const (
	TAG         = "orm"
	PLACEHOLDER = "?" // 占位符
)

type User struct {
	Id   int    `orm:"id"`
	Name string `orm:"name"`
	Age  int    `orm:"age"`
}

// makeWhere 构造 where 子句
func makeWhere(query []any) string {
	// 验证 where 子串是否是字符串
	whereStr, ok := query[0].(string)
	fmt.Println(whereStr)
	if !ok {
		panic("where 子串必须是 string")
	}
	// 验证占位符数量是否和传过来的参数数量一致
	placeHolderCount := strings.Count(whereStr, PLACEHOLDER)
	if placeHolderCount != len(query)-1 {
		panic("占位符数量和填充的参数数量不一致")
	}
	// 参数替换占位符
	for _, value := range query[1:] {
		t := reflect.TypeOf(value)
		var v string
		switch t.Kind() {
		case reflect.String:
			v = fmt.Sprintf("'%s'", value.(string))
		case reflect.Int:
			v = strconv.Itoa(value.(int))
		default:
			panic("unhandled default case")
		}
		fmt.Println(v)
		whereStr = strings.Replace(whereStr, PLACEHOLDER, v, 1)
	}
	return whereStr
}

func Find(obj any, query ...any) string {
	// 构造 where
	whereStr := makeWhere(query)

	t := reflect.TypeOf(obj)
	length := t.NumField()
	fields := make([]string, length)
	tableName := strings.ToLower(t.Name())
	for i := 0; i < length; i++ {
		field := t.Field(i)
		tag := field.Tag.Get(TAG)
		sqlField := fmt.Sprintf("`%s`", tag)
		fields[i] = sqlField

	}
	fieldStr := strings.Join(fields, ", ")
	fmt.Println(fieldStr)
	sqlStr := fmt.Sprintf("select %s from `%s` where %s", fieldStr, tableName, whereStr)
	return sqlStr

}

func main() {
	// select `id`, `name`, `age` from `user` where `name` = "name" and age > 10;
	sqlStr := Find(User{}, "name = ? and age > ?", "name", 10)
	fmt.Println(sqlStr)

}
