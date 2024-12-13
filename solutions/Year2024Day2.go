package solutions

import (
	"aoc/utils"
	"fmt"
	"math"
)

func Day2_2024_Part1() {
	path := "./inputs/Year2024Day2.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, line := range input {
		level := utils.StringSplitByChar(line, " ")
		if validLevel(level) {
			count += 1
		}
	}
	fmt.Println(count)
}

func Day2_2024_Part2() {
	path := "./inputs/Year2024Day2.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	for _, line := range input {
		level := utils.StringSplitByChar(line, " ")
		if validLevel(level) {
			count += 1
		} else {
			for i := 0; i < len(level); i++ {
				var newLevel []string
				for j := 0; j < len(level); j++ {
					if i != j {
						newLevel = append(newLevel, level[j])
					}
				}
				if validLevel(newLevel) {
					count += 1
					break
				}
			}
		}
	}
	fmt.Println(count)
}

func validLevel(level []string) bool {
	var diffs []int
	for i := 0; i < len(level)-1; i++ {
		num1, num2 := utils.ConvertStringToInt(level[i]), utils.ConvertStringToInt(level[i+1])
		diff := float64(num2 - num1)
		diffs = append(diffs, int(diff))
		if math.Abs(diff) < 1 || math.Abs(diff) > 3 {
			return false
		}
	}
	counts := []int{0, 0}
	for _, d := range diffs {
		if d < 0 {
			counts[0] += 1
		} else if d > 0 {
			counts[1] += 1
		}
	}
	if counts[0] > 0 && counts[1] > 0 {
		return false
	}
	return true
}
