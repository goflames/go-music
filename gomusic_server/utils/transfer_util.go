package utils

import (
	"strconv"
)

//将获取到的参数转为int8类型

func TransferToInt8(str string) int8 {
	targetInt64, _ := strconv.ParseInt(str, 10, 8)
	// 将 int64 转换为 int8
	target := int8(targetInt64)
	return target
}
