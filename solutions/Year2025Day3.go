package solutions

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Day3_2025_Part1() {
	path := "./inputs/Year2025Day3.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	res := 0
	for _, line := range input {
		strArr := strings.Split(line, "")
		start, end := 0, len(strArr)-2
		chars := make([]string, 0)
		for end < len(strArr) {
			localMax := -1
			indexUsed := 0
			for i := start; i <= end; i++ {
				num := utils.ConvertStringToInt(strArr[i])
				if num > localMax {
					localMax = num
					indexUsed = i
				}
			}
			start = indexUsed + 1
			end += 1
			chars = append(chars, utils.ConvertIntToString(localMax))
		}
		res += utils.ConvertStringToInt(strings.Join(chars, ""))
	}
	fmt.Println(res)
}

func Day3_2025_Part2() {
	path := "./inputs/Year2025Day3.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	res := 0
	for _, line := range input {
		strArr := strings.Split(line, "")
		start, end := 0, len(strArr)-12
		chars := make([]string, 0)
		for end < len(strArr) {
			localMax := -1
			indexUsed := 0
			for i := start; i <= end; i++ {
				num := utils.ConvertStringToInt(strArr[i])
				if num > localMax {
					localMax = num
					indexUsed = i
				}
			}
			start = indexUsed + 1
			end += 1
			chars = append(chars, utils.ConvertIntToString(localMax))
		}
		res += utils.ConvertStringToInt(strings.Join(chars, ""))
	}
	fmt.Println(res)
}
