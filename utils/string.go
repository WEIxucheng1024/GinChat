package utils

import (
	"math/rand"
	"time"
)

const code = "0123456789ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvxwyz-*"

// 生成字符串随机数
// size为字符串长度，max为随机数范围
func RandomString(size, max int) string {
	rand.Seed(time.Now().UnixNano())
	buffer := make([]byte, size, size)
	for i := 0; i < size; i++ {
		buffer[i] = code[rand.Intn(max)]
	}
	return string(buffer[:size])
}
