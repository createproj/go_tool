package sqltool

import (
	"container/list"
	"database/sql"
	"gotool/gotool/consts/comn_const"
	"reflect"
	"strings"

	"github.com/shopspring/decimal"
)

// Struct2selectSql 结构体对象的首层通过反射转换成sql及参数的list
func Struct2selectSql[T any](obj *T) (sqlReturn string) {
	sqlStr := ""
	if obj == nil {
		return sqlStr
	}
	typeOfObj := reflect.TypeOf(*obj)
	// 获取对象的反射值
	size := typeOfObj.NumField()
	for i := 0; i < size; i++ {
		dataDefine := typeOfObj.Field(i)
		tag := dataDefine.Tag
		gormTag := tag.Get("gorm")
		columnName, hasCol := strings.CutPrefix(gormTag, "column:")

		if tag == "" || !hasCol {
			continue
		}
		sqlStr = sqlStr + columnName + ", "
	}
	sqlStr = strings.TrimSpace(sqlStr)
	sqlStr, _ = strings.CutSuffix(sqlStr, comn_const.COMMA)
	sqlStr = strings.TrimSpace(sqlStr)
	return sqlStr
}

// Struct2whereSql 结构体对象的首层通过反射转换成sql及参数的list
func Struct2whereSql[T any](obj *T) (sqlReturn string, params list.List) {
	sqlStr := ""
	if obj == nil {
		return sqlStr, params
	}
	typeOfObj := reflect.TypeOf(*obj)
	// 获取对象的反射值
	objValue := reflect.ValueOf(*obj)
	size := objValue.NumField()
	for i := 0; i < size; i++ {
		dataDefine := typeOfObj.Field(i)
		valDef := objValue.Field(i)
		typeStr := dataDefine.Type.Name()
		tag := dataDefine.Tag
		gormTag := tag.Get("gorm")
		val := valDef.Interface()
		columnName, hasCol := strings.CutPrefix(gormTag, "column:")

		if tag == "" || !hasCol {
			continue
		}

		if strings.EqualFold(typeStr, "string") && val != nil {
			valStr := strings.TrimSpace(val.(string))
			if valStr != "" && gormTag != "" {
				sqlStr = sqlStr + " and " + columnName + "=?"
				params.PushBack(val)
			}
		} else if strings.EqualFold(typeStr, "int") && val != nil {
			valInt := val.(int)
			if valInt != 0 {
				sqlStr = sqlStr + " and " + columnName + "=?"
				params.PushBack(val)
			}
		} else if strings.EqualFold(typeStr, "int64") && val != nil {
			valInt64 := val.(int64)
			if valInt64 != 0 {
				sqlStr = sqlStr + " and " + columnName + "=?"
				params.PushBack(val)
			}
		} else if strings.EqualFold(typeStr, "Decimal") && val != nil {
			valDecimal := val.(decimal.Decimal)
			if valDecimal.String() != "0" {
				sqlStr = sqlStr + " and " + columnName + "=?"
				params.PushBack(val)
			}
		} else if strings.EqualFold(typeStr, "NullTime") && val != nil {
			valTime := val.(sql.NullTime)
			if valTime.Valid {
				sqlStr = sqlStr + " and " + columnName + "=?"
				params.PushBack(val)
			}
		}
	}
	return sqlStr, params
}

// PruneSql 修剪sql，去掉多余的and和前后空格
func PruneSql(sql string) string {
	if sql == "" || strings.TrimSpace(sql) == "" {
		return sql
	}
	sqlStr := strings.TrimSpace(sql)
	sqlStr = strings.ReplaceAll(sqlStr, comn_const.SQL_AND2, "")
	sqlStr, _ = strings.CutPrefix(sqlStr, comn_const.SQL_AND)
	sqlStr, _ = strings.CutSuffix(sqlStr, comn_const.SQL_AND)
	sqlStr = strings.TrimSpace(sqlStr)
	return sqlStr
}
