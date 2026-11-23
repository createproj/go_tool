package comn_const

import "time"

const (
	Y string = "Y"

	N       string = "N"
	W       string = "W"
	HEX_PRE string = "0x"

	SUCCESS string = "success"
	FAIL    string = "fail"
	WAIT    string = "wait"
	AFTER   string = "after"
	ROOT    string = "root"

	SQL_AND      string = "and"
	SQL_AND2     string = "and and"
	SQL_CONN_AND string = " and "

	NUM_STR0 string = "0"
	NUM_STR1 string = "1"
	NUM_STR2 string = "2"

	UINT64_0    uint64 = uint64(0)
	ZERO_PARENT string = "00000000000000000000"
	// eth钱包推理路径，默认前缀
	DEF_ETH_PATH_PRE string = "m/44'/60'/0'/0/"
	// TF_DAY 时间格式yyyyMMdd
	TF_DAY string = "20060102"
	FIX    string = "Fix"
	GRAB   string = "Grab"

	// BASE_KLINE_SCALE 基础K线刻度
	BASE_KLINE_SCALE = "1m"
)

var LocCN, _ = time.LoadLocation("Asia/Shanghai")
