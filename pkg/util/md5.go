package util

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 计算给定字符的MD5哈希值，返回其十六进制表示
func EncodeMD5(value string) string {
	//创建一个新的MD5计算器实例
	m := md5.New()

	//将value写入到MD5计算器中
	m.Write([]byte(value))

	//nil表示计算完哈希值后不添加后缀
	return hex.EncodeToString(m.Sum(nil))
}
