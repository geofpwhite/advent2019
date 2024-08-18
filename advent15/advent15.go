package main

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"

	adventhelperfunctions "github.com/geofpwhite/adventHelperFunctions"
	//	adventhelperfunctions "github.com/geofpwhite/adventHelperFunctions"
)

func main() {
	part1()
}

type node struct {
	value                 int
	up, down, left, right *node
}

var nullNode *node = &node{-1, nil, nil, nil, nil}

func parse() ([]int, int) {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
	// content, _ := os.ReadFile("test3.txt")
	// content, _ := os.ReadFile("test4.txt")
	lines := strings.Split(string(content), "\n")
	command := []int{}
	for _, numStr := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(strings.Replace(numStr, "\r", "", -1))
		if err != nil {
			fmt.Println(numStr, err)
		}
		command = append(command, num)
	}
	length := len(command)
	return command, length
}

type direction int

const NORTH direction = 1
const EAST direction = 2
const SOUTH direction = 3
const WEST direction = 4

type droid struct {
	x, y, xTry, yTry   int
	field              map[[2]int]int
	directionsTried    map[[2]int][]direction
	curDirectionTrying direction
}

func (r *droid) moveForward(dir direction) {
	switch dir {
	case NORTH:
		r.xTry, r.yTry = r.x, r.y
		r.yTry++
	case SOUTH:
		r.xTry, r.yTry = r.x, r.y
		r.yTry--
	case EAST:
		r.xTry, r.yTry = r.x, r.y
		r.xTry++
	case WEST:
		r.xTry, r.yTry = r.x, r.y
		r.xTry--
	}
}

func getCoordsForDirection(dir direction, x, y int) [2]int {
	switch dir {
	case NORTH:
		y++
	case SOUTH:
		y--
	case EAST:
		x++
	case WEST:
		x--
	}
	return [2]int{x, y}
}
func part1() {
	fmt.Println("f")
	base := 0
	command, length := parse()
	rob := droid{field: map[[2]int]int{{0, 0}: 1}, directionsTried: make(map[[2]int][]direction)}
	command = append(command, make([]int, 100000)...)
	cycle := rand.Intn(4) + 1
	// fmt.Println(command)
	defer func() {
		x := adventhelperfunctions.MapToImage[int](rob.field)
		filestr := ""
		for _, line := range x {
			for _, char := range line {
				if char == -1 {
					char = 3
				}
				filestr += strconv.Itoa(char)
			}
			filestr += "\n"
		}
		file, _ := os.Create("output.txt")
		file.WriteString(filestr)
		file.Close()
	}()

	for i := 0; i < length; {
		// println(i, command[i])
		numStr := strconv.Itoa(command[i])
		for len(numStr) < 5 {
			numStr = "0" + numStr
		}
		opcode, _ := strconv.Atoi(numStr[3:])
		paramModes := []byte{numStr[2], numStr[1], numStr[0]}
		// fmt.Println(string(paramModes[0]), string(paramModes[1]), string(paramModes[2]))
		var value1, value2 int
		if opcode == 99 {
			fmt.Println(opcode)
			break
		}

		if paramModes[0] == '1' {
			value1 = command[i+1]
		} else if paramModes[0] == '2' {
			value1 = command[base+command[i+1]]
		} else { //if opcode != 3 && opcode != 9 {

			value1 = command[command[i+1]]
		}
		if paramModes[1] == '1' {
			value2 = command[i+2]
		} else if paramModes[1] == '2' {
			value2 = command[base+command[i+2]]
		} else if opcode != 3 {
			value2 = command[command[i+2]]
		}
		// fmt.Println(opcode)
		switch opcode {
		case 1: //addition
			// println(numStr, command[i+3], string(paramModes[2]))
			// println(command[command[i+3]])
			if paramModes[2] == '2' {
				command[base+command[i+3]] = value1 + value2
			} else {
				command[command[i+3]] = value1 + value2
			}
			i += 4
		case 2: //multiplication
			if paramModes[2] == '2' {
				command[base+command[i+3]] = value1 * value2
			} else {
				command[command[i+3]] = value1 * value2
			}
			i += 4
		case 3: //input
			// if firstInput == -1 {
			// 	command[command[i+1]] = secondInput
			// } else {
			// 	command[command[i+1]] = firstInput
			// 	firstInput = -1
			// }
			for j := cycle; j%5 > 0; j = (j + 1) % 5 {
				if !slices.Contains(rob.directionsTried[[2]int{rob.x, rob.y}], direction(j)) && rob.field[getCoordsForDirection(direction(j), rob.x, rob.y)] != -1 {
					rob.curDirectionTrying = direction(j)
					fmt.Println(direction(j), rob.x, rob.y)
					rob.moveForward(rob.curDirectionTrying)
					break
				}
			}
			cycle = (cycle % 4) + 1
			if paramModes[0] == '2' {
				command[base+command[i+1]] = int(rob.curDirectionTrying)
			} else {
				command[command[i+1]] = int(rob.curDirectionTrying)
			}

			i += 2
		case 4: //output

			if rob.x > 1000 || rob.y > 1000 ||
				rob.x < -1000 || rob.y < -1000 {
				println("break")
				return
			}

			if value1 == 1 {
				rob.field[[2]int{rob.xTry, rob.yTry}] = 1
				rob.directionsTried[[2]int{rob.x, rob.y}] = append(rob.directionsTried[[2]int{rob.x, rob.y}], rob.curDirectionTrying)
				rob.x, rob.y = rob.xTry, rob.yTry
			} else if value1 == 2 {
				rob.field[[2]int{rob.xTry, rob.yTry}] = 2
				rob.x, rob.y = rob.xTry, rob.yTry
				return

			} else {
				rob.field[[2]int{rob.xTry, rob.yTry}] = -1
				rob.directionsTried[[2]int{rob.x, rob.y}] = append(rob.directionsTried[[2]int{rob.x, rob.y}], rob.curDirectionTrying)
			}
			// fmt.Println(value1, rob.curDirectionTrying)
			// fmt.Println(numStr, " -> ", value1)
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(command[i+1]))
			// fmt.Println(output, " output ", i, opcode)
			i += 2
		case 5:
			if value1 != 0 {
				i = value2
			} else {
				i += 3
			}
		case 6:
			if value1 == 0 {
				i = value2
			} else {
				i += 3
			}
		case 7:
			if value1 < value2 {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 1
				} else {
					command[command[i+3]] = 1
				}
			} else {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 0
				} else {
					command[command[i+3]] = 0
				}
			}
			i += 4
		case 8:
			if value1 == value2 {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 1
				} else {
					command[command[i+3]] = 1
				}
			} else {
				if paramModes[2] == '2' {
					command[base+command[i+3]] = 0
				} else {
					command[command[i+3]] = 0
				}
			}
			i += 4
		case 9:
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(command[i+1]))
			// fmt.Printf("%d,%d\n", value1, command[i+1])
			fmt.Printf("%d\n", base)
			base += value1
			// base += command[i+1]
			// fmt.Printf("%d\n", base)
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(value1))
			i += 2
		case 99:
			break
		}

	}
	fmt.Println(rob.field)

}
