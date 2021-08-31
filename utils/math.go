package utils

import (
	"strconv"
	"strings"
)

var (
	Tokens string
	Length int
)

func init() {
	// 0-9
	for i := 0; i <= 9; i++ {
		Tokens += strconv.Itoa(i)
	}

	// a-z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('a') + byte(i))
	}
	// A-Z
	for i := 0; i < 26; i++ {
		Tokens += string(byte('A') + byte(i))
	}
	Length = len(Tokens)
}

// IdToString id转62进制
func IdToString(id int) string {
	var res string
	for id > 0 {
		d := id % Length
		res = string(Tokens[d]) + res
		id /= Length
	}
	return res
}

func StringToId(str string) int {
	var res = 0
	for _, s := range str {
		value := strings.Index(Tokens, string(s))
		res = res*Length + value
	}
	return res
}
