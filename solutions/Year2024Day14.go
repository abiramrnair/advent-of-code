package solutions

import (
	"aoc/utils"
	"fmt"
	"time"
)

func Day14_2024_Part1() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day14.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ymod,xmod := 103,101
	m := make(map[[2]int]int)
	for _,line := range input {
		l := utils.StringSplitByChar(line," ")
		pos,vel :=  utils.StringSplitByChar(l[0],"p="),utils.StringSplitByChar(l[1],"v=")
		p := utils.StringSplitByChar(pos[1],",")
		v := utils.StringSplitByChar(vel[1],",")
		px,py := utils.ConvertStringToInt(p[0]),utils.ConvertStringToInt(p[1])
		vx,vy := utils.ConvertStringToInt(v[0]),utils.ConvertStringToInt(v[1])
		nx := (px + (vx * 100)) % xmod
		ny := (py + (vy * 100)) % ymod
		if nx < 0 {
			nx = xmod + nx
		}
		if ny < 0 {
			ny = ymod + ny
		}
		_,ok := m[[2]int{ny,nx}]
		if !ok {
			m[[2]int{ny,nx}] = 0
		}
		m[[2]int{ny,nx}] += 1
	}
	q1,q2,q3,q4 := 0,0,0,0
	for k,z := range m {
		cy,cx := k[0],k[1]
		if cy >= 0 && cy < 51 && cx >= 0 && cx < 50 {
			q1 += z
		}
		if cy > 51 && cy <= 102 && cx >= 0 && cx < 50 {
			q2 += z
		}
		if cy >= 0 && cy < 51 && cx > 50 && cx <= 100 {
			q3 += z
		}
		if cy > 51 && cy <= 102 && cx > 50 && cx <= 100 {
			q4 += z
		}
	}
	answer := q1 * q2 * q3 * q4
	fmt.Println("Answer:", answer)
}
// 169
func Day14_2024_Part2() {
	defer utils.CodeTimer()()
	INPUT_PATH := "./inputs/Year2024Day14.txt"
	input, err := utils.GetInputAsArrayOfStrings(INPUT_PATH)
	if err != nil {
		fmt.Println(err)
	}
	ymod,xmod := 103,101
	// Image was found on 7753 for input.
	i := 0
	for i < 10000 {
		m := make(map[[2]int]int)
		for _,line := range input {
			l := utils.StringSplitByChar(line," ")
			pos,vel :=  utils.StringSplitByChar(l[0],"p="),utils.StringSplitByChar(l[1],"v=")
			p := utils.StringSplitByChar(pos[1],",")
			v := utils.StringSplitByChar(vel[1],",")
			px,py := utils.ConvertStringToInt(p[0]),utils.ConvertStringToInt(p[1])
			vx,vy := utils.ConvertStringToInt(v[0]),utils.ConvertStringToInt(v[1])
			nx := (px + (vx * i)) % xmod
			ny := (py + (vy * i)) % ymod
			if nx < 0 {
				nx = xmod + nx
			}
			if ny < 0 {
				ny = ymod + ny
			}
			_,ok := m[[2]int{ny,nx}]
			if !ok {
				m[[2]int{ny,nx}] = 0
			}
			m[[2]int{ny,nx}] += 1
		}
		grid := make([][]string,0)
		for i := 0; i < 103; i++ {
			row := make([]string,0)
			for j := 0; j < 101; j++ {
				row = append(row, " ")
			}
			grid = append(grid, row)
		}
		for k,_ := range m {
			cy,cx := k[0],k[1]
			grid[cy][cx] = "*"
		}
		utils.Print2DArray(grid)
		fmt.Println("Seconds:", i)
		time.Sleep(10 * time.Millisecond)
		i += 1
	}
}
