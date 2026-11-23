package objtool

import (
	"bytes"
	"container/list"
	"database/sql"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"gotool/gotool/consts/comn_const"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Workiva/go-datastructures/set"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/shopspring/decimal"
)

// TrueFunc 如果condition为真则返回originalId，否则调用falseApply方法并以params作为参数
func TrueFunc(condition bool, originalId interface{}, falseApply func(*[]string) interface{}, params *[]string) interface{} {
	if condition {
		return originalId
	}
	return falseApply(params)
}

// TrueStr 如果condition为真则返回trueVal，否则返回falseVal
func TrueStr(condition bool, trueVal, falseVal string) string {
	if condition {
		return trueVal
	}
	return falseVal
}

// TrueInt32 如果condition为真则返回trueVal，否则返回falseVal
func TrueInt32(condition bool, trueVal, falseVal int32) int32 {
	if condition {
		return trueVal
	}
	return falseVal
}

// TrueInt 如果condition为真则返回trueVal，否则返回falseVal
func TrueInt(condition bool, trueVal, falseVal int) int {
	if condition {
		return trueVal
	}
	return falseVal
}

// TrueInt64 如果condition为真则返回trueVal，否则返回falseVal
func TrueInt64(condition bool, trueVal, falseVal int64) int64 {
	if condition {
		return trueVal
	}
	return falseVal
}

// TrueFuncInt64 如果condition为真则返回originalId，否则调用falseApply方法并以params作为参数
func TrueFuncInt64(condition bool, originalId int64, falseApply func(string) int64, params string) int64 {
	if condition {
		return originalId
	}
	return falseApply(params)
}

func TrueFunc1Int64(condition bool, originalId int64, falseApply func() int64) int64 {
	if condition {
		return originalId
	}
	return falseApply()
}

func SyncMap2Map(syncMap *sync.Map) map[string]interface{} {
	m := make(map[string]interface{})
	// 遍历sync.Map
	syncMap.Range(func(key, value interface{}) bool {
		// 将sync.Map中的键值对复制到普通map中
		m[key.(string)] = value
		return true // 返回true继续遍历
	})
	return m
}

func DeciStr(deci *decimal.Decimal) string {
	if deci == nil {
		return comn_const.NUM_STR0
	}
	return deci.String()
}

func NoNilArr[T any](arr []T) bool {
	return arr != nil && len(arr) > 0
}

func AsNilArr[T any](arr []T) bool {
	return arr == nil || len(arr) == 0
}

// Arr2IntfArr 任意对象数组转成interface对象数组
func Arr2IntfArr[T any](arr []T) []interface{} {
	if AsNilArr(arr) {
		return nil
	}
	intfArr := make([]interface{}, len(arr))
	for i, v := range arr {
		intfArr[i] = v
	}
	return intfArr
}

// Intf2ObjArr interface对象数组转成任意对象数组
func Intf2ObjArr[T any](arr []interface{}) []T {
	if AsNilArr(arr) {
		return nil
	}
	intfArr := make([]T, len(arr))
	for i, v := range arr {
		intfArr[i] = v.(T)
	}
	return intfArr
}

const layoutDate = "2006-01-02"
const layoutDateTime = "2006-01-02 15:04:05"
const layoutDateTimeMilli = "2006-01-02 15:04:05.000"

func ParseTime2UtcStr(timeObj time.Time) string {
	formatted := timeObj.Format("2006-01-02T15:04:05.000Z")
	return formatted
}
func ParseSqlTime2UtcStr(timeObj sql.NullTime) string {
	if !timeObj.Valid {
		return ""
	}
	return ParseTime2UtcStr(timeObj.Time)
}
func ParseSqlTime2int64(timeObj sql.NullTime) int64 {
	if !timeObj.Valid {
		return 0
	}
	return timeObj.Time.UnixMilli()
}

// Parse1Date2Milli 时间格式： 2006-01-02
func Parse1Date2Milli(dataStr1 string) int64 {
	// 使用Parse函数解析时间字符串
	parsedTime, err := time.Parse(layoutDate, dataStr1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}
	return parsedTime.UnixMilli()
}

func Parse1DateTime2Milli(dataStr1 string) int64 {
	// 使用Parse函数解析时间字符串
	parsedTime, err := time.Parse(layoutDateTime, dataStr1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}
	return parsedTime.UnixMilli()
}

// Parse1Date2NullTime 格式： 2006-01-02
func Parse1Date2NullTime(dataStr1 string) *sql.NullTime {
	// 使用Parse函数解析时间字符串
	parsedTime, err := time.Parse(layoutDate, dataStr1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}
	return &sql.NullTime{Time: parsedTime, Valid: true}
}

// Int64NullTime 毫秒格式： 1668873660000
func Int64NullTime(dataStr1 int64) *sql.NullTime {
	return &sql.NullTime{Time: time.UnixMilli(dataStr1), Valid: true}
}

// Str64NullTime 毫秒格式： 1668873660000
func Str64NullTime(dataStr1 string) *sql.NullTime {
	return &sql.NullTime{Time: time.UnixMilli(Str2int64(dataStr1)), Valid: true}
}

func JsonNumber2NullTime(timeNumber json.Number) *sql.NullTime {
	if timeNumber == "" {
		return nil
	}
	val := sql.NullTime{Time: time.UnixMilli(gconv.Int64(timeNumber.String())), Valid: true}
	return &val
}

func CurrNullTime() *sql.NullTime {
	return &sql.NullTime{Time: time.Now(), Valid: true}
}

func Parse1DateTime2NullTime(dataStr1 string) *sql.NullTime {
	// 使用Parse函数解析时间字符串
	parsedTime, err := time.Parse(layoutDateTime, dataStr1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}
	return &sql.NullTime{Time: parsedTime, Valid: true}
}

func Parse1DateTimeMilli2NullTime(dataStr1 string) *sql.NullTime {
	// 使用Parse函数解析时间字符串
	parsedTime, err := time.Parse(layoutDateTimeMilli, dataStr1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil
	}
	return &sql.NullTime{Time: parsedTime, Valid: true}
}
func ParseInt64NullTime(ts int64) *sql.NullTime {
	if ts == 0 {
		return &sql.NullTime{Time: time.UnixMilli(0)}
	}
	return &sql.NullTime{Time: time.UnixMilli(ts), Valid: true}
}

func Parse1DateTimeMilli2int64(dataStr1 string) int64 {
	// 使用Parse函数解析时间字符串
	parsedTime, err := time.Parse(layoutDateTimeMilli, dataStr1)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return 0
	}
	return parsedTime.UnixMilli()
}

// DateMilli2Str1 时间戳转UTC时间。输出格式： 2006-01-02 15:04:05.000
func DateMilli2Str1(milli int64) string {
	// 使用Parse函数解析时间字符串
	val := time.UnixMilli(milli).UTC().Format(layoutDateTimeMilli)
	return val
}

// Milli2Local 时间戳转本地时间。输出格式： 2006-01-02 15:04:05.000
func Milli2Local(milli int64) string {
	// 使用Parse函数解析时间字符串
	val := time.UnixMilli(milli).Format(layoutDateTimeMilli)
	return val
}

// DateStr 输出格式： 2006-01-02 15:04:05.000
func DateStr(t time.Time) string {
	// 使用Parse函数解析时间字符串
	val := t.UTC().Format(layoutDateTimeMilli)
	return val
}
func CurrStr() string {
	// 使用Parse函数解析时间字符串
	val := time.Now().UTC().Format(layoutDateTimeMilli)
	return val
}

func JsonNumber2decimal(timeNumber json.Number) *decimal.Decimal {
	if timeNumber == "" {
		return nil
	}
	val := decimal.RequireFromString(timeNumber.String())
	return &val
}

func SpendMilli(t1 *time.Time) int64 {
	if t1 == nil {
		return 0
	}
	return time.Now().UnixMilli() - t1.UnixMilli()
}

func SonicToInt64(v interface{}) int64 {
	switch t := v.(type) {
	case string:
		n, _ := strconv.ParseInt(t, 10, 64)
		return n
	case float64:
		return int64(t)
	case json.Number:
		n, _ := t.Int64()
		return n
	default:
		return 0
	}
}

func SonicToDecimal(v interface{}) decimal.Decimal {
	switch t := v.(type) {
	case string:
		return decimal.RequireFromString(t)
	case float32:
		return decimal.NewFromFloat32(t)
	case float64:
		return decimal.RequireFromString(strconv.FormatFloat(t, 'f', 8, 64))
	case json.Number:
		return *(JsonNumber2decimal(t))
	default:
		return decimal.Decimal{}
	}
}

// ArrObjVal 取数组对象中的属性的值，返回新数组
func ArrObjVal[T any](arr []T, getPropVal func(obj T) interface{}) []interface{} {
	if AsNilArr(arr) {
		return nil
	}
	intfArr := make([]interface{}, len(arr))
	for i, v := range arr {
		intfArr[i] = getPropVal(v)
	}
	return intfArr
}

func List2array[T any](list *list.List) []T {
	if list == nil || list.Len() == 0 {
		return nil
	}

	size := list.Len()
	arrayIdx := 0
	arr := make([]T, size)
	for e := list.Front(); e != nil; e = e.Next() {
		el := e.Value.(T)
		arr[arrayIdx] = el
		arrayIdx++
	}
	return arr
}

func SetMapValSet[B any](mapPoint *map[string]*set.Set, keyObj string, val B) {
	if mapPoint == nil {
		return
	}
	mapObj := *mapPoint
	setObj := mapObj[keyObj]
	if setObj == nil {
		setObj = set.New()
		mapObj[keyObj] = setObj
	}
	setObj.Add(val)
}

func List2set[T any](list *list.List) *set.Set {
	if list == nil || list.Len() == 0 {
		return nil
	}
	arrayIdx := 0
	setObj := set.New()
	for e := list.Front(); e != nil; e = e.Next() {
		el := e.Value.(T)
		setObj.Add(el)
		arrayIdx++
	}
	return setObj
}

func Array2list[T any](array []T) list.List {
	lis := list.List{}
	if array == nil || len(array) == 0 {
		return lis
	}

	for i, size := 0, len(array); i < size; i++ {
		lis.PushBack(array[i])
	}
	return lis
}

// Array2set 数组转成集合
func Array2set(arr []any) *set.Set {
	addressSet := set.New()
	for _, v := range arr {
		//TODO	可能要添加更多的类型匹配
		switch v.(type) {
		case string:
			addressSet.Add(v.(string))
		case []uint8:
			addressSet.Add(string(v.([]uint8)))
		default:
			addressSet.Add(v.(string))
		}

	}
	return addressSet
}

// Set2array 集合转成数组
func Set2array[T any](setObj *set.Set) []T {
	arr := make([]T, setObj.Len())
	if setObj == nil || setObj.Len() == 0 {
		return arr
	}
	intfArr := setObj.Flatten()
	for i, v := range intfArr {
		arr[i] = v.(T)
	}
	return arr
}

func TrueFirst[T any](condition bool, first, second T) T {
	if condition {
		return first
	} else {
		return second
	}
}

func ReverseArray[T any](arr []T) []T {
	if arr == nil || len(arr) == 0 {
		return arr
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func StrCommaSplit2int64(originalStr string) []int64 {
	return StrSplit2int64(originalStr, comn_const.COMMA)
}

func Str2int64(str string) int64 {
	intVal, _ := strconv.ParseInt(str, 10, 64)
	return intVal
}
func Str2int(str string) int {
	intVal, _ := strconv.Atoi(str)
	return intVal
}

func StrSplit2int64(originalStr, splitStr string) []int64 {
	if originalStr == "" {
		return nil
	}

	if splitStr == "" {
		originalStr = ","
	}
	strArr := strings.Split(originalStr, splitStr)
	listObj := list.List{}

	for _, v := range strArr {
		intVal, _ := strconv.ParseInt(v, 10, 64)
		listObj.PushBack(intVal)
	}
	return List2array[int64](&listObj)
}

func GapMilli(t1, t2 time.Time) int64 {
	return t2.Sub(t1).Milliseconds()
}

// IntToBytes 数字转字节数组
/*
参考： https://blog.csdn.net/qq_45535340/article/details/124955107
*/
func IntToBytes(n int64) []byte {
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, n)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int64 {
	byteBuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(byteBuff, binary.BigEndian, &data)
	return int64(data)
}

// GetRelativeDifference 计算 arr1 相对于 arr2 的相对差集
func GetRelativeDifference(arr1, arr2 []string) []string {
	if len(arr1) == 0 {
		return arr2
	}
	// 创建一个 map 用于记录 arr2 中的元素
	elementMap := make(map[string]bool)
	for _, element := range arr2 {
		elementMap[element] = true
	}

	var result []string
	// 遍历 arr1，将不在 arr2 中的元素添加到结果切片中
	for _, element := range arr1 {
		if !elementMap[element] {
			result = append(result, element)
		}
	}
	return result
}
