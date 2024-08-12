package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	part2()
}

type tile int

const EMPTY, WALL, BLOCK, HOR_PADDLE, BALL tile = 0, 1, 2, 3, 5

func part1() (int, error) {
	base := 0
	screen := map[[2]int]tile{}
	curTileCoords := [2]int{}
	mod3 := 0
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
				command[base+command[i+1]] = 2
			} else {
				command[command[i+1]] = 2
			}

			i += 2
		case 4: //output

			output += strconv.Itoa(value1)
			switch mod3 {
			case 0:
				curTileCoords[0] = value1
			case 1:
				curTileCoords[1] = value1
			case 2:
				screen[curTileCoords] = tile(value1)
			}
			mod3 = (mod3 + 1) % 3
			fmt.Println(value1)
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
	blocks := 0
	for _, val := range screen {
		if val == BLOCK {
			blocks++
		}
	}
	fmt.Println(blocks, screen)
	// fmt.Println(command, "f")
	if err != nil {
		// panic("problem, output is " + output)
		return -1, fmt.Errorf("bad output somehow")
	}
	return outputNum, nil

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

func part2() (int, error) {
	base := 0
	screen := map[[2]int]tile{}
	curTileCoords := [2]int{}
	mod3 := 0
	command, _ := parse()
	command[0] = 2
	command = append(command, make([]int, 100000)...)
	// fmt.Println(command)
	output := ""
	paddlePosition, ballPosition := [2]int{}, [2]int{}
	for i := 0; ; { //i < length; {
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
			joystick := 0
			if ballPosition[0] > paddlePosition[0] {
				joystick = 1
			} else if ballPosition[0] < paddlePosition[0] {
				joystick = -1
			}
			if paramModes[0] == '2' {
				command[base+command[i+1]] = joystick
			} else {
				command[command[i+1]] = joystick
			}

			i += 2
		case 4: //output

			switch mod3 {
			case 0:
				curTileCoords[0] = value1
			case 1:
				curTileCoords[1] = value1
			case 2:
				if curTileCoords == [2]int{-1, 0} {
					fmt.Println("score", value1)
				} else {
					screen[curTileCoords] = tile(value1)
					if value1 == 3 {
						paddlePosition[0], paddlePosition[1] = curTileCoords[0], curTileCoords[1]
					}
					if value1 == 4 {
						ballPosition[0], ballPosition[1] = curTileCoords[0], curTileCoords[1]
					}
				}
			}
			mod3 = (mod3 + 1) % 3
			fmt.Println(value1)
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
	blocks := 0
	realScreen := [50][50]tile{}
	for key, val := range screen {
		realScreen[key[0]][key[1]] = val
		if val == BLOCK {
			blocks++
		}
	}
	for _, line := range realScreen {
		fmt.Println(line)
	}
	// fmt.Println(command, "f")
	if err != nil {
		// panic("problem, output is " + output)
		return -1, fmt.Errorf("bad output somehow")
	}
	return outputNum, nil

}
