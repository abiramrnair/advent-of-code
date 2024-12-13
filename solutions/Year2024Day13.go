package solutions

import (
	"aoc/utils"
	"fmt"
)

type Machine struct {
	A    utils.Coord
	B    utils.Coord
	Goal utils.Coord
}

func Day13_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day13.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	machines := parseInput(input)
	tokens := 0
	for _, m := range machines {
		tokens += getMachineTokens(m)
	}
	fmt.Println("Answer:", tokens)
}

func Day13_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day13.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	machines := parseInput(input)
	tokens := 0
	for _, m := range machines {
		tokens += getMachineTokens2(m)
	}
	fmt.Println("Answer:", tokens)
}

func getMachineTokens(m Machine) int {
	ax, ay := m.A.Row, m.A.Col
	bx, by := m.B.Row, m.B.Col
	gx, gy := m.Goal.Row, m.Goal.Col
	axmult, bxmult := ax*ay, bx*ay
	_, bymult := ay*ax, by*ax
	gxmult, gymult := gx*ay, gy*ax
	btokens := float64(gxmult-gymult) / float64(bxmult-bymult)
	if !utils.IsWholeNumber(btokens) {
		return 0
	}
	atokens := float64(gxmult-(bxmult*int(btokens))) / float64(axmult)
	if !utils.IsWholeNumber(atokens) {
		return 0
	}
	return (int(atokens) * 3) + int(btokens)
}

func getMachineTokens2(m Machine) int {
	ax, ay := m.A.Row, m.A.Col
	bx, by := m.B.Row, m.B.Col
	gx, gy := m.Goal.Row+10000000000000, m.Goal.Col+10000000000000
	axmult, bxmult := ax*ay, bx*ay
	_, bymult := ay*ax, by*ax
	gxmult, gymult := gx*ay, gy*ax
	btokens := float64(gxmult-gymult) / float64(bxmult-bymult)
	if !utils.IsWholeNumber(btokens) {
		return 0
	}
	atokens := float64(gxmult-(bxmult*int(btokens))) / float64(axmult)
	if !utils.IsWholeNumber(atokens) {
		return 0
	}
	return (int(atokens) * 3) + int(btokens)
}

func parseInput(input []string) []Machine {
	var machines []Machine
	l := 0
	for l < len(input) {
		m := Machine{}
		lineA, lineB, lineC := input[l], input[l+1], input[l+2]
		ax := utils.StringSplitByChar(lineA, "X+")
		ay := utils.StringSplitByChar(lineA, "Y+")
		ax = utils.StringSplitByChar(ax[1], ",")
		m.A.Row = utils.ConvertStringToInt(ax[0])
		m.A.Col = utils.ConvertStringToInt(ay[len(ay)-1])
		bx := utils.StringSplitByChar(lineB, "X+")
		by := utils.StringSplitByChar(lineB, "Y+")
		bx = utils.StringSplitByChar(bx[1], ",")
		m.B.Row = utils.ConvertStringToInt(bx[0])
		m.B.Col = utils.ConvertStringToInt(by[len(by)-1])
		px := utils.StringSplitByChar(lineC, "X=")
		px = utils.StringSplitByChar(px[1], ",")
		py := utils.StringSplitByChar(lineC, "Y=")
		m.Goal.Row = utils.ConvertStringToInt(px[0])
		m.Goal.Col = utils.ConvertStringToInt(py[len(py)-1])
		machines = append(machines, m)
		l += 4
	}
	return machines
}
