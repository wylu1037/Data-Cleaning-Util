package util

import "strconv"

// Str2uint64 字符串转uint64数字
func Str2uint64(str string) uint64 {
	num, _ := strconv.Atoi(str)
	return uint64(num)
}
