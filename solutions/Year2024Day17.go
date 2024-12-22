package solutions

import (
	"aoc/utils"
	"fmt"
)

func Day17_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day17.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	computer := Day173BitComputer{}
	computer.Init(input)
	computer.ProcessInstructions()
	computer.PrintState()
}

func Day17_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day17.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	computer := Day173BitComputer{}
	computer.Init(input)
	var match string
	for _,n := range computer.Instructions {
		match += utils.ConvertIntToString(n)
	}
	for i := 0; i < 100000000000; i++ {
		computer := Day173BitComputer{}
		computer.Init2(input,i)
		computer.ProcessInstructions()
		output := computer.GetOutput()
		if output == "3,1,7,5,5,3,0" || output == "1,7,5,5,3,0" || output == "7,5,5,3,0" || output == "5,5,3,0" || output == "5,3,0" || output == "3,0" || output == "0" {
			fmt.Println(output,":",i)
		}
	}
}

type Day173BitComputer struct {
	A,B,C,InstructionPointer int
	Instructions []int
	Output []int
}

func (c *Day173BitComputer) Init(input []string) {
	registerALine := utils.StringTrimSpaces(utils.StringSplitByChar(input[0],":")[1])
	registerBLine := utils.StringTrimSpaces(utils.StringSplitByChar(input[1],":")[1])
	registerCLine := utils.StringTrimSpaces(utils.StringSplitByChar(input[2],":")[1])
	programLine := utils.StringSplitByChar(utils.StringTrimSpaces(utils.StringSplitByChar(input[4],":")[1]),",")
	c.A = utils.ConvertStringToInt(registerALine)
	c.B = utils.ConvertStringToInt(registerBLine)
	c.C = utils.ConvertStringToInt(registerCLine)
	c.InstructionPointer = 0
	c.Instructions = make([]int,0)
	c.Output = make([]int,0)
	for _,p := range programLine {
		c.Instructions = append(c.Instructions, utils.ConvertStringToInt(p))
	}
}

func (c *Day173BitComputer) Init2(input []string, aVal int) {
	// registerALine := utils.StringTrimSpaces(utils.StringSplitByChar(input[0],":")[1])
	registerBLine := utils.StringTrimSpaces(utils.StringSplitByChar(input[1],":")[1])
	registerCLine := utils.StringTrimSpaces(utils.StringSplitByChar(input[2],":")[1])
	programLine := utils.StringSplitByChar(utils.StringTrimSpaces(utils.StringSplitByChar(input[4],":")[1]),",")
	c.A = aVal
	c.B = utils.ConvertStringToInt(registerBLine)
	c.C = utils.ConvertStringToInt(registerCLine)
	c.InstructionPointer = 0
	c.Instructions = make([]int,0)
	c.Output = make([]int,0)
	for _,p := range programLine {
		c.Instructions = append(c.Instructions, utils.ConvertStringToInt(p))
	}
}

func (c Day173BitComputer) GetComboOperand(n int) int {
	if n < 4 {
		return n
	} else if n == 4 {
		return c.A
	} else if n == 5 {
		return c.B
	} else if n == 6 {
		return c.C
	}
	return 0
}

func (c *Day173BitComputer) ProcessInstructions() {
	for c.InstructionPointer < len(c.Instructions) {
		opcode,operand := c.Instructions[c.InstructionPointer],c.Instructions[c.InstructionPointer + 1]
		if opcode == 0 {
			numerator := c.A
			denominator := utils.PowInt(2, c.GetComboOperand(operand))
			div := int(numerator / denominator)
			c.A = div
			c.InstructionPointer += 2
		} else if opcode == 1 {
			c.B = c.B ^ operand
			c.InstructionPointer += 2
		} else if opcode == 2 {
			c.B =  c.GetComboOperand(operand) % 8
			c.InstructionPointer += 2
		} else if opcode == 3 {
			if c.A != 0 {
				c.InstructionPointer = operand
			} else {
				c.InstructionPointer += 2
			}
		} else if opcode == 4 {
			c.B = c.B ^ c.C
			c.InstructionPointer += 2
		} else if opcode == 5 {
			c.Output = append(c.Output, c.GetComboOperand(operand) % 8)
			c.InstructionPointer += 2
		} else if opcode == 6 {
			numerator := c.A
			denominator := utils.PowInt(2, c.GetComboOperand(operand))
			div := int(numerator / denominator)
			c.B = div
			c.InstructionPointer += 2
		} else if opcode == 7 {
			numerator := c.A
			denominator := utils.PowInt(2, c.GetComboOperand(operand))
			div := int(numerator / denominator)
			c.C = div
			c.InstructionPointer += 2
		}
	}
}

func (c Day173BitComputer) GetOutput() string {
	var res []string
	for _,n := range c.Output {
		res = append(res, utils.ConvertIntToString(n))
	}
	return utils.StringJoinByChar(res,",")
}

func (c Day173BitComputer) PrintState() {
	fmt.Println("Register A:", c.A)
	fmt.Println("Register B:", c.B)
	fmt.Println("Register C:", c.C)
	fmt.Println("Output:", c.GetOutput())
}