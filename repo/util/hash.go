package util

import (
	"encoding/binary"
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
func To62Str(i uint32) string {
	sb := strings.Builder{}
	for i > 0 {
		sb.WriteString(c62[i%62])
		i /= 62
	}
	return sb.String()
}

// Murmur3 hash算法
func Murmur3(key []byte) uint32 {
	const (
		c1 = 0xcc9e2d51
		c2 = 0x1b873593
		r1 = 15
		r2 = 13
		m  = 5
		n  = 0xe6546b64
	)
	var (
		seed = uint32(1938)
		h    = seed
		k    uint32
		l    = len(key)
		end  = l - (l % 4)
	)
	for i := 0; i < end; i += 4 {
		k = binary.LittleEndian.Uint32(key[i:])
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2

		h ^= k
		h = (h << r2) | (h >> (32 - r2))
		h = h*m + n
	}
	k = 1
	switch l & 3 {
	case 3:
		k ^= uint32(key[end+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[end+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[end])
		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2
		h ^= k
	}
	h ^= uint32(l)
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	h ^= h >> 16
	return h
}
