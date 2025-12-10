package solutions

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func Day2_2025_Part1() {
	path := "./inputs/Year2025Day2.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	input = utils.StringSplitByChar(input[0], ",")
	invalidIds := 0
	for _, ranges := range input {
		r := utils.StringSplitByChar(ranges, "-")
		start := utils.ConvertStringToInt(r[0])
		end := utils.ConvertStringToInt(r[1])
		for i := start; i < end+1; i++ {
			number := utils.ConvertIntToString(i)
			if len(number)%2 == 0 {
				firstHalf := number[:len(number)/2]
				secondHalf := number[len(number)/2:]
				if firstHalf == secondHalf {
					invalidIds += utils.ConvertStringToInt(number)
				}
			}
		}
	}
	fmt.Println(invalidIds)
}

func Day2_2025_Part2() {
	path := "./inputs/Year2025Day2.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	input = utils.StringSplitByChar(input[0], ",")
	invalidIds := 0
	for _, ranges := range input {
		r := utils.StringSplitByChar(ranges, "-")
		start := utils.ConvertStringToInt(r[0])
		end := utils.ConvertStringToInt(r[1])
		for i := start; i < end+1; i++ {
			number := utils.ConvertIntToString(i)
			numberLength := len(number)
			chosenNumber := 0
			for j := 1; j < numberLength; j++ {
				if numberLength%j == 0 {
					arr := make([]string, 0)
					for z := 0; z < len(number); z += j {
						subNumber := number[z : z+j]
						arr = append(arr, subNumber)
					}
					if len(arr) >= 2 {
						isValid := true
						for k := 1; k < len(arr); k++ {
							first := arr[k-1]
							second := arr[k]
							if first != second {
								isValid = false
								break
							}
						}
						if isValid {
							invalidId := utils.ConvertStringToInt(strings.Join(arr, ""))
							if invalidId != chosenNumber {
								invalidIds += invalidId
							}
							chosenNumber = invalidId
						}
					}
				}
			}
		}
	}
	fmt.Println(invalidIds)
}
