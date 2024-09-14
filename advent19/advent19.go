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

type coords [2]int

func neighbors(c coords, field map[coords]bool) []coords {
	f1, f2 := field[coords{c[0], c[1] + 1}], field[coords{c[0] + 1, c[1]}]
	switch f1 && f2 {
	case true:
		return []coords{}
	default:
		switch f1 || f2 {
		case false:
			return []coords{
				{c[0] + 1, c[1]},
				{c[0], c[1] + 1},
			}
		default:
			if field[coords{c[0], c[1] + 1}] {
				return []coords{
					{c[0] + 1, c[1]},
				}
			} else {
				return []coords{
					{c[0], c[1] + 1},
				}
			}
		}

	}
}

func part1() {
	base := 0
	command, length := parse()
	field := map[coords]bool{}
	command = append(command, make([]int, 500000)...)
	queue := []coords{{950, 1220}}
	stack := []coords{}
	check := 'q'
	curX, curY := 300, 110
	xChosen := false
	defer func() {
		fmt.Println(curX-99, curY, curX)
		fmt.Println(field[coords{curX, curY}])
		fmt.Println(field[coords{curX - 99, curY}])
		fmt.Println(field[coords{curX - 99, curY + 99}])
		fmt.Println(field[coords{curX, curY + 99}])
		// curX--
		curY--
		fmt.Println(field[coords{curX, curY}])
		fmt.Println(field[coords{curX - 99, curY}])
		fmt.Println(field[coords{curX - 99, curY + 99}])
		fmt.Println(field[coords{curX, curY + 99}])

		for field[coords{curX, curY + 99}] && field[coords{curX - 99, curY + 99}] && field[coords{curX - 99, curY}] && field[coords{curX, curY}] {
			curX--
			curY--
		}

	}()
	// fmt.Println(command)
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
		// if opcode == 99 {
		// 	fmt.Println(opcode)
		// 	break
		// }

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
				if xChosen {
					command[base+command[i+1]] = curY
					xChosen = false
				} else {
					command[base+command[i+1]] = curX
					xChosen = true
				}
			} else {
				if xChosen {
					command[command[i+1]] = curY
					xChosen = false
				} else {
					command[command[i+1]] = curX
					xChosen = true
				}
			}

			i += 2
		case 4: //output

			c := coords{curX, curY}
			if value1 == 0 && check != 'q' {
				stack = []coords{}
			}
			if value1 == 1 && check == '3' {
				field[c] = true
				return
			}
			if value1 == 1 && check == '2' {
				field[c] = true
			}
			if value1 == 1 && check == '1' {
				field[c] = true
			}

			fmt.Println(c)
			if !field[c] && value1 == 1 && check == 'q' {
				field[c] = true

				queue = append(queue, (neighbors(coords{curX, curY}, field))...)
				stack = append(stack, coords{curX + 99, curY}, coords{curX, curY + 99}, coords{curX + 99, curY + 99})
			}
			if value1 == 0 && len(queue) == 0 {
				queue = append(queue, (neighbors(coords{curX, curY}, field))...)
			}
			length := len(stack)
			// fmt.Println(string(check), length, len(queue))

			switch length {
			case 0:
				curX, curY = queue[0][0], queue[0][1]
				fmt.Println(len(queue))
				queue = queue[1:]
				fmt.Println(len(queue))
				check = 'q'
			case 1:
				curX, curY = stack[0][0], stack[0][1]
				stack = []coords{}
				check = '3'
			case 2:
				curX, curY = stack[1][0], stack[1][1]
				stack = stack[:1]
				check = '2'
				// fmt.Println(len(stack))
			case 3:
				curX, curY = stack[2][0], stack[2][1]
				stack = stack[:2]
				check = '1'
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
			i, base = 0, 0
			command, _ = parse()
			command = append(command, make([]int, 500000)...)
		}

	}

	fmt.Println(len(field))
	realField := make([]string, 500)
	for i := range realField {

		for j := i - i/3; j < i*4/3; j++ {
			if j > 500 {
				break
			}
			if field[coords{i, j}] {
				realField[i] += "#"
			} else {
				realField[i] += "."
			}
		}
	}
	fileStr := strings.Join(realField, "\n")
	file, _ := os.Create("output.txt")
	file.WriteString(fileStr)
	file.Close()

}
