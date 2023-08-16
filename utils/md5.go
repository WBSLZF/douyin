package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// 输入数据并返回MD5哈希值的16进制字符串表示作为函数返回结果
func Md5Encode(data string) string {
	h := md5.New()                     //创建一个MD5哈希对象
	h.Write([]byte(data))              //将字节数组写入MD5哈希对象
	tempStr := h.Sum(nil)              //计算并返回MD5哈希对象的哈希值，参数nil表示不添加任何额外的数据
	return hex.EncodeToString(tempStr) //将MD5的哈希值的16进制字符串表示函数返回的结果
}

// 加密
func MakePassWord(plainwd string) string {
	return Md5Encode(plainwd)
}

// 判断是否相等，若不相等则返回true
func ValidPassWord(plainwd, password string) bool {
	encode_password := Md5Encode(plainwd)
	return encode_password != password
}
