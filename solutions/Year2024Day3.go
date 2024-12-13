package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day3_2024_Part1() {
	path := "./inputs/Year2024Day3.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for _, line := range input {
		nomuls := utils.StringSplitByChar(line, "mul")
		for _, nomul := range nomuls {
			idxopen := utils.StringFirstIndexOfChar(nomul, "(")
			idxclosed := utils.StringFirstIndexOfChar(nomul, ")")
			idxcomma := utils.StringFirstIndexOfChar(nomul, ",")
			if (string(nomul[0]) != "(") ||
				(idxopen == -1 || idxclosed == -1 || idxcomma == -1) ||
				(idxcomma-idxopen > 4 || idxclosed-idxcomma > 4) {
				continue
			}
			firstNum := ""
			secondNum := ""
			i := idxopen + 1
			for i < idxcomma {
				firstNum += string(nomul[i])
				i += 1
			}
			i = idxcomma + 1
			for i < idxclosed {
				secondNum += string(nomul[i])
				i += 1
			}
			sum += utils.ConvertStringToInt(firstNum) * utils.ConvertStringToInt(secondNum)
		}
	}
	fmt.Println(sum)
}

func Day3_2024_Part2() {
	path := "./inputs/Year2024Day3.txt"
	input, err := utils.GetInputAsArrayOfStrings(path)
	if err != nil {
		fmt.Println(err)
	}
	sum := 0
	for _, line := range input {
		nomuls := utils.StringSplitByChar(line, "mul")
		enabled := true
		for _, nomul := range nomuls {
			idxopen := utils.StringFirstIndexOfChar(nomul, "(")
			idxclosed := utils.StringFirstIndexOfChar(nomul, ")")
			idxcomma := utils.StringFirstIndexOfChar(nomul, ",")
			if (string(nomul[0]) != "(") ||
				(idxopen == -1 || idxclosed == -1 || idxcomma == -1) ||
				(idxcomma-idxopen > 4 || idxclosed-idxcomma > 4) {
				if utils.StringHasChar(nomul, "do()") {
					enabled = true
				}
				if utils.StringHasChar(nomul, "don't()") {
					enabled = false
				}
				continue
			}
			firstNum := ""
			secondNum := ""
			i := idxopen + 1
			for i < idxcomma {
				firstNum += string(nomul[i])
				i += 1
			}
			i = idxcomma + 1
			for i < idxclosed {
				secondNum += string(nomul[i])
				i += 1
			}
			if enabled {
				sum += utils.ConvertStringToInt(firstNum) * utils.ConvertStringToInt(secondNum)
			}
			if utils.StringHasChar(nomul, "do()") {
				enabled = true
			}
			if utils.StringHasChar(nomul, "don't()") {
				enabled = false
			}
		}
	}
	fmt.Println(sum)
}
