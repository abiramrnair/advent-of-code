package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day11_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day11.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	parsed := utils.StringSplitByChar(input[0], " ")
	m := make(map[string]int)
	for _, block := range parsed {
		_, ok := m[block]
		if !ok {
			m[block] = 1
		} else {
			m[block] += 1
		}
	}
	b := 0
	for b < 25 {
		m = transformMap(m)
		b += 1
	}
	answer := 0
	for _, v := range m {
		answer += v
	}
	fmt.Println("Answer:", answer)
}

func Day11_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day11.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	parsed := utils.StringSplitByChar(input[0], " ")
	m := make(map[string]int)
	for _, block := range parsed {
		_, ok := m[block]
		if !ok {
			m[block] = 1
		} else {
			m[block] += 1
		}
	}
	b := 0
	for b < 75 {
		m = transformMap(m)
		b += 1
	}
	answer := 0
	for _, v := range m {
		answer += v
	}
	fmt.Println("Answer:", answer)
}

func transformMap(m map[string]int) map[string]int {
	nm := make(map[string]int)
	for key, value := range m {
		if key == "0" {
			_, ok := nm["1"]
			if !ok {
				nm["1"] = 0
			}
			nm["1"] += value
		} else if len(key)%2 == 0 {
			leftSide := trimLeadingZeros(key[0 : len(key)/2])
			rightSide := trimLeadingZeros(key[len(key)/2:])
			_, ok := nm[leftSide]
			if !ok {
				nm[leftSide] = value
			} else {
				nm[leftSide] += value
			}
			_, ok = nm[rightSide]
			if !ok {
				nm[rightSide] = value
			} else {
				nm[rightSide] += value
			}
		} else {
			newKey := utils.ConvertIntToString(utils.ConvertStringToInt(key) * 2024)
			_, ok := nm[newKey]
			if !ok {
				nm[newKey] = value
			} else {
				nm[newKey] += value
			}
		}
	}
	return nm
}

/** Inefficient **/
// func transformString(s string) string {
// 	var res []string
// 	stringArr := utils.StringSplitByChar(s, " ")
// 	for _,block := range stringArr {
// 		intblock := utils.ConvertStringToUint(block)
// 		if intblock == 0 {
// 			res = append(res, "1")
// 		} else if len(block) % 2 == 0 {
// 			leftSide := trimLeadingZeros(block[0:len(block)/2])
// 			rightSide := trimLeadingZeros(block[len(block)/2:])
// 			res = append(res, leftSide)
// 			res = append(res, rightSide)
// 		} else {
// 			res = append(res, utils.ConvertUintToString(intblock * 2024))
// 		}
// 	}
// 	return utils.StringJoinByChar(res, " ")
// }

func trimLeadingZeros(s string) string {
	stringArr := utils.StringSplitByChar(s, "")
	for i := 0; i < len(stringArr); i++ {
		if stringArr[i] == "0" && i != len(stringArr)-1 {
			stringArr[i] = ""
		} else {
			break
		}
	}
	return utils.StringJoinByChar(stringArr, "")
}
