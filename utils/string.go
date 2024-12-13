package utils

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func StringHasChar(str string, substr string) bool {
	return strings.Contains(str, substr)
}

func StringSplitByChar(str string, char string) []string {
	return strings.Split(str, char)
}

func StringJoinByChar(str []string, char string) string {
	return strings.Join(str, char)
}

func StringTrimSpaces(str string) string {
	return strings.TrimSpace(str)
}

func ConvertStringToInt(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func ConvertIntToString(num int) string {
	return strconv.Itoa(num)
}

func ConvertStringToUint(str string) uint64 {
	num, _ := strconv.ParseUint(str, 10, 64)
	return num
}

func ConvertUintToString(num uint64) string {
	res := fmt.Sprintf("%v", num)
	return res
}

func StringFirstIndexOfChar(str string, char string) int {
	return strings.Index(str, char)
}

func StringIsLowercase(str string) bool {
	for _, r := range str {
		if !unicode.IsLower(r) {
			return false
		}
	}
	return true
}

func StringIsUppercase(str string) bool {
	for _, r := range str {
		if !unicode.IsUpper(r) {
			return false
		}
	}
	return true
}

func ConvertStringToLowercase(str string) string {
	return strings.ToLower(str)
}

func ConverStringToUppercase(str string) string {
	return strings.ToUpper(str)
}
