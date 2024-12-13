package utils

import (
	"fmt"
	"math"
	"time"
)

func PrintArray(array []string) {
	for _, line := range array {
		fmt.Println(line)
	}
}

func Print2DArray(array [][]string) {
	for _, line := range array {
		fmt.Println(StringJoinByChar(line, ""))
	}
}

func CodeTimer() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Runtime: %v\n", time.Since(start))
	}
}

func IsWholeNumber(x float64) bool {
	return math.Ceil(x) == x
}
