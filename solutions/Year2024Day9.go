package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day9_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day9.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	inputVals := utils.StringSplitByChar(input[0], "")
	systemArr := make([]string, 0)
	id := 0
	for i, val := range inputVals {
		if i%2 == 0 {
			j := 0
			for j < utils.ConvertStringToInt(val) {
				systemArr = append(systemArr, utils.ConvertIntToString(id))
				j += 1
			}
			id += 1
		} else {
			j := 0
			for j < utils.ConvertStringToInt(val) {
				systemArr = append(systemArr, ".")
				j += 1
			}
		}
	}
	i, j := 0, len(systemArr)-1
	for i < j {
		for systemArr[i] != "." {
			i += 1
		}
		for systemArr[j] == "." {
			j -= 1
		}
		if i < j {
			tmp := systemArr[i]
			systemArr[i] = systemArr[j]
			systemArr[j] = tmp
		}
	}
	res := getSystemChecksum(systemArr)
	fmt.Println(res)
}

func Day9_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day9.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	inputVals := utils.StringSplitByChar(input[0], "")
	systemArr := make([]string, 0)
	id := 0
	for i, val := range inputVals {
		if i%2 == 0 {
			j := 0
			for j < utils.ConvertStringToInt(val) {
				systemArr = append(systemArr, utils.ConvertIntToString(id))
				j += 1
			}
			id += 1
		} else {
			j := 0
			for j < utils.ConvertStringToInt(val) {
				systemArr = append(systemArr, ".")
				j += 1
			}
		}
	}
	j := len(systemArr) - 1
	for j >= 0 {
		for j >= 0 && systemArr[j] == "." {
			j -= 1
		}
		jstart := j
		for jstart-1 >= 0 && systemArr[jstart-1] == systemArr[j] {
			jstart -= 1
		}
		i := 0
		for i < len(systemArr) {
			for i < len(systemArr) && systemArr[i] != "." {
				i += 1
			}
			iend := i
			for iend+1 < len(systemArr) && systemArr[iend+1] == systemArr[i] {
				iend += 1
			}
			if iend < jstart && iend-i >= j-jstart {
				z := 0
				for z < (j-jstart)+1 {
					tmp := systemArr[i+z]
					systemArr[i+z] = systemArr[jstart+z]
					systemArr[jstart+z] = tmp
					z += 1
				}
				break
			}
			i = iend + 1
		}
		j = jstart - 1
	}
	res := getSystemChecksum(systemArr)
	fmt.Println(res)
}

func getSystemChecksum(arr []string) int {
	res := 0
	for i, val := range arr {
		res += (i * utils.ConvertStringToInt(val))
	}
	return res
}
