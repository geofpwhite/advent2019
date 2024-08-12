package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
}

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

const NORTH direction = 0
const EAST direction = 1
const SOUTH direction = 2
const WEST direction = 3

type robot struct {
	x, y int
	dir  direction
}

func (r *robot) turnRight() {
	r.dir = (r.dir + 1) % 4
}
func (r *robot) turnLeft() {
	r.dir = (r.dir + 3) % 4
}
func (r *robot) moveForward() {
	switch r.dir {
	case NORTH:
		r.y++
	case SOUTH:
		r.y--
	case EAST:
		r.x++
	case WEST:
		r.x--
	}
}

func part1() (int, error) {
	base := 0
	rob := new(robot)
	whiteCoords := make(map[[2]int]bool)
	whiteCoords[[2]int{0, 0}] = true
	painted := make(map[[2]int]bool)
	command, length := parse()
	command = append(command, make([]int, 100000)...)
	// fmt.Println(command)
	output := ""
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
			if paramModes[0] == '2' {
				if whiteCoords[[2]int{rob.x, rob.y}] {
					command[base+command[i+1]] = 1
				} else {
					command[base+command[i+1]] = 0
				}
			} else {
				if whiteCoords[[2]int{rob.x, rob.y}] {
					command[command[i+1]] = 1
				} else {
					command[command[i+1]] = 0
				}
			}

			i += 2
		case 4: //output
			output += strconv.Itoa(value1)
			fmt.Println(output, " output")
			if len(output) == 2 {
				if value1 == 0 {
					rob.turnLeft()
				} else {
					rob.turnRight()
				}
				rob.moveForward()
				output = ""
			} else {
				painted[[2]int{rob.x, rob.y}] = true
				fmt.Println(true)
				if value1 == 0 {
					whiteCoords[[2]int{rob.x, rob.y}] = false
				} else {
					whiteCoords[[2]int{rob.x, rob.y}] = true
				}
			}
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
			// fmt.Printf("%d\n", base)
			base += value1
			// base += command[i+1]
			// fmt.Printf("%d\n", base)
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(value1))
			i += 2
		case 99:
			break
		}

	}
	outputNum, err := strconv.Atoi(output)
	fmt.Println(outputNum)
	fmt.Println(len(painted))
	fmt.Println((painted))
	lines := [100]string{}
	line := ""
	for i := 0; i < 100; i++ {
		line += "."
	}
	for i := range lines {
		lines[i] = line
	}
	for key, val := range whiteCoords {
		if val {
			lines[key[0]+50] = lines[key[0]+50][:key[1]+50] + "#" + lines[key[0]+50][key[1]+51:]
		}
	}
	for i := range lines {
		fmt.Println(lines[i])
	}
	fmt.Println(lines)
	// fmt.Println(command, "f")
	if err != nil {
		// panic("problem, output is " + output)
		return -1, fmt.Errorf("bad output somehow")
	}
	return outputNum, nil

}
