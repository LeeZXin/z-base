package util

import (
	"github.com/spaolacci/murmur3"
	"strings"
)

var (
	c62 = []string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}
)

// To62Str 数字转化为62进制字符串
func To62Str(i uint64) string {
	sb := strings.Builder{}
	for i > 0 {
		sb.WriteString(c62[i%62])
		i /= 62
	}
	return sb.String()
}

// Murmur3 hash算法
func Murmur3(key []byte) uint64 {
	if key == nil {
		return 0
	}
	h64 := murmur3.New64WithSeed(2)
	h64.Write(key)
	return h64.Sum64()
}
