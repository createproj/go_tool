package strtool

import (
	"context"
	"github.com/bytedance/sonic"
	"github.com/gogf/gf/v2/os/glog"
)

func SonicStrByPath(byteArray []byte, path ...interface{}) string {
	return SonicValByPath(false, byteArray, path...)
}

func SonicValByPath(asRow bool, byteArray []byte, path ...interface{}) string {
	nodeVal, nodeErr := sonic.Get(byteArray, path...)
	if nodeErr != nil {
		return ""
	}

	if asRow {
		val, err := nodeVal.Raw()
		if err != nil {
			return ""
		}
		return val
	} else {
		val, err := nodeVal.String()
		if err != nil {
			return ""
		}
		return val
	}
}

func Sonic2array2(ctx context.Context, message string) [][]interface{} {
	if IsBlank(message) {
		return nil
	}
	var items [][]interface{}
	if err := sonic.Unmarshal([]byte(message), &items); err != nil {
		glog.Infof(ctx, "sonic.Unmarshal data -> [][]interface{} error: %v", err)
		return nil
	}
	return items
}
