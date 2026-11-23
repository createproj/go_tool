package strtool

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"gotool/gotool/consts/comn_const"
	"gotool/gotool/util/objtool"
	"math"
	"math/big"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/Workiva/go-datastructures/set"
	"github.com/bytedance/sonic"
	strutil2 "github.com/duke-git/lancet/v2/strutil"
	jsoniter "github.com/json-iterator/go"
	"github.com/shopspring/decimal"
)

func EasyPwd() string {
	currTime, _ := strconv.ParseInt(time.Now().Format("2006010215"), 10, 64)
	currTimeStr := strconv.FormatInt(currTime*2+1, 10)
	return currTimeStr
}

func Asy(tag string) bool {
	return strings.EqualFold(tag, comn_const.Y)
}
func Noy(tag string) bool {
	return !strings.EqualFold(tag, comn_const.Y)
}

func EqualAnyIgnoreCase(originalStr *string, judgeStrs ...string) bool {
	if originalStr == nil {
		return false
	}

	for _, v := range judgeStrs {
		if strings.EqualFold(*originalStr, v) {
			return true
		}
	}
	return false
}

// EqualForce 强制比较字符串是否相等。不忽略大小写。sql中检查是否有大写字母：SELECT * FROM wd_spread_balance WHERE up_address ~ '[A-Z]' limit 99
func EqualForce(aa, bb string) bool {
	return strings.Compare(aa, bb) == 0
}

func NoEqualForce(aa, bb string) bool {
	return strings.Compare(aa, bb) != 0
}

func Int64TransStr(num int64) string {
	val := strconv.FormatInt(num, 10)
	return val
}

func Int32TransStr(num int32) string {
	val := strconv.FormatInt(int64(num), 10)
	return val
}

func IntTransStr(num int) string {
	val := strconv.Itoa(num)
	return val
}

func Str2int64(num string) int64 {
	val, _ := strconv.ParseInt(num, 10, 64)
	return val
}
func Str2int(num string) int {
	val, _ := strconv.Atoi(num)
	return val
}
func Str2Float32(num string) float32 {
	val, _ := strconv.ParseFloat(num, 32)
	return float32(val)
}
func SplitCommaStr2Set(strs *string) *set.Set {
	return SplitStr2Set(strs, comn_const.COMMA)
}
func SplitCommaStr2SetArr(strs *string) []string {
	return objtool.Set2array[string](SplitStr2Set(strs, comn_const.COMMA))
}

// Sanitize 去掉多余的特殊字符
func Sanitize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "\u200b", "") // 零宽空格
	s = strings.ReplaceAll(s, "\ufeff", "") // BOM
	return s
}

// DumpString 打印字节（十六进制）和 rune（索引 + 字符显示 + U+XXXX）
func DumpString(s string) {
	fmt.Printf("====== (len bytes=%d, len runes=%d) ===\n", len(s), utf8.RuneCountInString(s))
	fmt.Printf("bytes: % x\n", []byte(s))
	i := 0
	for pos, r := range s {
		fmt.Printf("%03d @byte%02d: rune=%q U+%04X\n", i, pos, r, r)
		i++
	}
}

func IsEmpty(s string) bool {
	return s == ""
}

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == "" || IsEmpty(s)
}

func HasBlank(strs ...string) bool {
	if len(strs) == 0 {
		return false
	}
	for _, str := range strs {
		if IsBlank(str) {
			return true
		}
	}
	return false
}
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

func Nil2Blank(str *string) string {
	if str == nil {
		return ""
	}
	return *str
}

func ValNoBlank(strs ...string) string {
	for _, str := range strs {
		if !IsBlank(str) {
			return str
		}
	}
	return ""
}

// RepairStrRepeat 修复字符串中重复的
func RepairStrRepeat(str, splitStr string, ignoreStr *set.Set) {

}

func SplitStr2Set(strs *string, splitStr string) *set.Set {
	setObj := set.New()
	if strs == nil || *strs == "" {
		return setObj
	}
	strArr := strings.Split(*strs, splitStr)
	for _, v := range strArr {
		if v == "" {
			continue
		}
		setObj.Add(v)
	}
	return setObj
}

// Obj2JsonStr 对象转json字符串
func Obj2JsonStr(obj interface{}) string {
	json, _ := jsoniter.MarshalToString(obj)
	return json
}

func Json2Obj[T any](jsonStr string) T {
	var obj T
	err := sonic.Unmarshal([]byte(jsonStr), &obj)
	if err != nil {
		fmt.Sprintf("Json2ObjFail=%v err=%v", jsonStr, err)
	}
	return obj
}

func Json2ArrObj[T any](jsonStr string) []T {
	var objs []T
	err := sonic.Unmarshal([]byte(jsonStr), &objs)
	if err != nil {
		fmt.Sprintf("Json2ObjFail=%v err=%v", jsonStr, err)
	}
	return objs
}

// HexToBigFloat 16进制转10进制，并且除掉小数
/**
 * @param hex: 16进制的数字
 * @param decimalQty: 16进制数字的小数位数
 */
func HexToBigFloat(hex string, decimalQty int) big.Float {
	hexInt := big.Int{}
	hexInt.SetString(hex, 16)

	floatNum, divFloat := big.Float{}, big.Float{}
	floatNum.SetInt(&hexInt)

	// num, _, err1 := big.ParseFloat(hexInt, 16, uint(prec), big.ToNearestAway)
	decimalCalc := strconv.FormatFloat(math.Pow(float64(10), float64(decimalQty)), 'f', -1, 64)
	divFloat.SetString(decimalCalc)

	floatNumResult := floatNum.Quo(&floatNum, &divFloat)
	return *floatNumResult
}

// HexToInt64 16进制转int64数字
func HexToInt64(hex string) int64 {
	hex, _ = strings.CutPrefix(hex, comn_const.HEX_PRE)
	decNum, _ := strconv.ParseInt(hex, 16, 64)
	return decNum
}

// HexToBigInt 16进制转bigInt
func HexToBigInt(hex string) *big.Int {
	hex, _ = strings.CutPrefix(hex, comn_const.HEX_PRE)

	hexInt := big.Int{}
	hexInt.SetString(hex, 16)

	//deci, _ := decimal.NewFromString(hex)
	//var retBig = deci.BigInt()
	//retBig.SetString(hex, 16)
	return &hexInt
}

// HexByte2str 16进制转字符串
func HexByte2str(hexStr string) string {
	hexStr, _ = strings.CutPrefix(hexStr, comn_const.HEX_PRE)
	buf := make([]byte, len(hexStr)/2)
	n, _ := hex.Decode(buf, []byte(hexStr))
	byteArr := buf[:n]
	return string(byteArr)
}

// Str2ByteHex 字符串先按utf8转成字节数组，再把字节数组转成16进制字符串
func Str2ByteHex(str string) string {
	byteArr := strutil2.StringToBytes(str)
	return hex.EncodeToString(byteArr)
}

// MnemonicFormat 助记词格式化
func MnemonicFormat(mnemonicStr string) string {
	mnemonicStr = strings.ReplaceAll(mnemonicStr, comn_const.LineBreak, comn_const.SPACE)
	mnemonicStr = strings.ReplaceAll(mnemonicStr, comn_const.TAB, comn_const.SPACE)
	mnemonicStr = strings.TrimSpace(mnemonicStr)
	words := strings.Split(mnemonicStr, comn_const.SPACE)
	aimStr := ""
	for i, v := range words {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		if i == 0 {
			aimStr = v
		} else {
			aimStr = aimStr + comn_const.SPACE + v
		}
	}
	return aimStr
}

// ValConvert
/**
 * 值转换。如果没有匹配的值，返回默认值，否则返回。参数中的数组，第1，3，5...个是匹配值，第2，4，6...个是返回值，
 * @param matchVal 判断值
 * @param defVal 默认值
 * @param candidate 比对值数组
 * @return 判断后符合的值
 */
func ValConvert(matchVal, defVal string, candidate ...string) string {
	if len(candidate) < 2 {
		return defVal
	}
	for i, length := 0, len(candidate)-1; i < length; i += 2 {
		if matchVal == candidate[i] {
			if i+1 > length-1 {
				return defVal
			}
			return candidate[i+1]
		}
	}
	return defVal
}

func CutPre(str, prefix string) string {
	aim, _ := strings.CutPrefix(str, prefix)
	return aim
}

func HexToDate(hex string) time.Time {
	timeStamp := HexToInt64(hex)
	tMs := time.Unix(timeStamp, 0)
	return tMs
}

// 数学相关

func TrimTail0(num string) string {
	s, _ := strings.CutSuffix(num, comn_const.NUM_STR0)
	return s
}

// FormatDeci 格式化小数，去掉多余的0
func FormatDeci(num *decimal.Decimal) string {
	if num == nil || num.IsZero() {
		return comn_const.ZERO
	}
	return num.String()
}

func StrDef0(num string) decimal.Decimal {
	param, err := decimal.NewFromString(num)
	if err != nil {
		param = decimal.NewFromInt(0)
	}
	return param
}

// StrNegative 转为负责
func StrNegative(num string) string {
	param := StrDef0(num)
	if param.GreaterThan(decimal.NewFromInt(0)) {
		param = param.Mul(decimal.NewFromInt(-1))
	}
	return param.String()
}

// StrAbs 字符串数字转绝对值
func StrAbs(num string) string {
	param := StrDef0(num)
	if param.LessThan(decimal.NewFromInt(0)) {
		param = param.Mul(decimal.NewFromInt(-1))
		return param.String()
	}
	return num
}

func DefaultStrZero(str *interface{}, newStr string) {
	if str == nil {
		*str = "0"
	} else {
		*str = newStr
	}
}

// StrMin 字符串取最小值
func StrMin(aa, bb string) string {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	if aaDeci.LessThan(bbDeci) {
		return aa
	}
	return bb
}

// StrGreat 字符串数字大于
func StrGreat(aa, bb string) bool {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.GreaterThan(bbDeci)
}

// StrGreatEqual 字符串数字大于等于
func StrGreatEqual(aa, bb string) bool {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.GreaterThanOrEqual(bbDeci)
}

// StrCmpEqual 字符串 运算等于
func StrCmpEqual(aa, bb string) bool {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.Equal(bbDeci)
}

var ZeroDeci = decimal.NewFromInt(0)

// StrLessEqual0 字符串数字小于
func StrLessEqual0(aa string) bool {
	aaDeci := StrDef0(aa)
	return aaDeci.LessThanOrEqual(ZeroDeci)
}

// StrLess 字符串数字小于
func StrLess(aa, bb string) bool {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.LessThan(bbDeci)
}

// StrLessEqual 字符串数字小于等于
func StrLessEqual(aa, bb string) bool {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.LessThanOrEqual(bbDeci)
}

// StrAdd 字符串数字相加
func StrAdd(aa, bb string) string {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.Add(bbDeci).String()
}

// StrMinus 字符串数字相减
func StrMinus(aa, bb string) string {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.Sub(bbDeci).String()
}

// StrMul 字符串数字相乘
func StrMul(aa, bb string) string {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	return aaDeci.Mul(bbDeci).String()
}

// StrDivDown4 字符串数字相除
func StrDivDown4(aa, bb string) string {
	return StrDivDown(aa, bb, 4)
}

// StrDivDown 字符串数字相除
func StrDivDown(aa, bb string, deci int) string {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	val := aaDeci.DivRound(bbDeci, int32(deci)).String()
	val = strings.TrimRight(val, "0")
	val = strings.TrimRight(val, ".")
	return val
}

// StrDiv 字符串数字相除
func StrDiv(aa, bb string) string {
	aaDeci := StrDef0(aa)
	bbDeci := StrDef0(bb)
	val := aaDeci.DivRound(bbDeci, 18).String()
	val = strings.TrimRight(val, "0")
	val = strings.TrimRight(val, ".")
	return val
}

// StrDown2 字符串数字保留2位小数，多的数字舍掉
func StrDown2(aa string) string {
	aaDeci := StrDef0(aa)
	return aaDeci.RoundDown(2).String()
}

// StrDown3 字符串数字保留2位小数，多的数字舍掉
func StrDown3(aa string) string {
	aaDeci := StrDef0(aa)
	return aaDeci.RoundDown(3).String()
}

// StrDown6 字符串数字保留2位小数，多的数字舍掉
func StrDown6(aa string) string {
	aaDeci := StrDef0(aa)
	return aaDeci.RoundDown(6).String()
}

// SuperSplit 超级分割 切记!这个分割去重
func SuperSplit(addresses string) ([]string, int) {
	addresses = strings.ToLower(addresses)
	addresses = replaceAllSymbol(addresses)
	//清理掉为空的数据
	step1 := strings.Split(addresses, comn_const.VerticalLine)
	var res []any
	objSet := set.New()
	for _, v := range step1 {
		if !strings.EqualFold(v, "") {
			objSet.Add(v)
		}
	}
	res = objSet.Flatten()
	lenOfArr := len(res)
	return anyArr2StrArr(res), lenOfArr
}

// FormatStrNum 格式化字符串数字,保留小数位
func FormatStrNum(num string, deci int) string {

	num1, _ := decimal.NewFromString(num)
	trimmedNum1 := num1.Truncate(int32(deci))

	return trimmedNum1.String()
}

// anyArr2StrArr any数组转string数组 传入any类型数组,返回str类型数组
func anyArr2StrArr(anyArr []any) []string {
	res := make([]string, len(anyArr))
	for i, v := range anyArr {
		res[i] = v.(string)
	}
	return res
}

// rplaceAllSymbol 替换所有符号为"|"
func replaceAllSymbol(str string) string {
	oldSymbol := []string{
		comn_const.COMMA,
		comn_const.LineBreak,
		comn_const.SPACE,
		comn_const.SEMICOLON,
		comn_const.SemicolonZh,
		comn_const.DASHED,
		comn_const.XG,
		comn_const.CommaZh}
	//遍历oldSymbol
	for _, v := range oldSymbol {
		str = strings.ReplaceAll(str, v, comn_const.VerticalLine)
	}
	return str
}

// DatePlusDays 日期计算 单位 天
func DatePlusDays(start string, change int) (sql.NullTime, sql.NullTime) {
	dateLayout := "2006-01-02"
	date, err := time.Parse(dateLayout, start)
	dateEnd := date.Add(time.Duration(change) * time.Hour * 24)
	if err != nil {
		return sql.NullTime{}, sql.NullTime{}
	}
	return sql.NullTime{date, true}, sql.NullTime{dateEnd, true}
}

// SuperSplit 超级分割
func SplitToInterFace(addresses string, symbol string) []interface{} {
	step1 := strings.Split(addresses, symbol)
	var res []interface{}
	for _, v := range step1 {
		if !strings.EqualFold(v, "") {
			res = append(res, v)
		}
	}
	return res
}

var reg1 = regexp.MustCompile(".{1,64}")

// 按长度分割，举例: "123456789" 分割为 ["123","456","789"]
func SplitGroupByLength(hexWithout0x string, eachLength int) []string {
	var str = CutPre(hexWithout0x, "0x")
	var arr = reg1.FindAllString(str, 1000000)
	return arr
}

func String2Int64(intStr string) int64 {
	num, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func String2Dec(intStr string) decimal.Decimal {
	num, err := decimal.NewFromString(intStr)
	if err != nil {
		return decimal.Zero
	}
	return num
}

// 系统级的字符串工具

func CapturePanic(err interface{}) string {
	if err == nil {
		return ""
	}

	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip 3 层：Callers, CapturePanic, defer func

	frames := runtime.CallersFrames(pcs[:n])

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("panic: %v\n", err))
	sb.WriteString("call stack:\n")

	for {
		frame, more := frames.Next()
		sb.WriteString(fmt.Sprintf("    %s\n        %s:%d\n",
			frame.Function, frame.File, frame.Line))
		if !more {
			break
		}
	}

	return sb.String()
}
