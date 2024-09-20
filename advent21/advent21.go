package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func part1() {
	command, length := parse()
	base := 0
	command = append(command, make([]int, 100000)...)
	output := ""
	// var WALK = []byte("WALK")
	var WALK, walkStr = 0, ""
	for _, b := range []byte("WALK\n") {
		walkStr += strconv.Itoa(int(b))
	}
	WALK, _ = strconv.Atoi(walkStr)
	var NEWLINE = 10
	var AND, andStr = 0, ""
	for _, b := range []byte("AND") {
		andStr += strconv.Itoa(int(b))
	}
	AND, _ = strconv.Atoi(andStr)
	var OR, orStr = 0, ""
	for _, b := range []byte("OR") {
		orStr += strconv.Itoa(int(b))
	}
	OR, _ = strconv.Atoi(orStr)
	var NOT, notStr = 0, ""
	for _, b := range []byte("NOT") {
		notStr += strconv.Itoa(int(b))
	}
	SPACE := strconv.Itoa(int(' '))
	NOT, _ = strconv.Atoi(notStr)
	var A, B, C, D, T, J = strconv.Itoa(int('A')), strconv.Itoa(int('B')), strconv.Itoa(int('C')), strconv.Itoa(int('D')), strconv.Itoa(int('T')), strconv.Itoa(int('J'))
	fmt.Println(A, B, C, D, T, J, AND, OR, NOT, andStr, orStr, notStr, walkStr, SPACE, WALK, NEWLINE)
	not_d_j_walk := 0
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
			// notdj := notStr + SPACE + D + SPACE + J + "10"
			// str := "NOT C T\nOR T J\nNOT A T\nOR T J\nNOT B T\n OR T J\nAND D J\nWALK\n"
			str := "NOT I T\n NOT T T\nAND E T\nOR H T\nNOT T J\nNOT J J\nAND D J\nNOT B T\nAND T J\nNOT A T\nOR T J\nRUN\n"
			// horei := "NOT I T\n NOT T T\nAND E T\nOR H T\nOR T J\n"
			//jump if  (not a or not b or not c) and (d) and ((h) or (e and i))
			//without parentheses is
			//NOT E T\nNOT T T\nAND I T\nOR H T\nAND D T\nNOT A J\nOR
			//(e and i) or h ) and d and (not a or not b or not c)
			//
			//
			// NOT H T\nNOT T T\nNOT T J\nNOT E T\nNOT T T\nOR I T\nOR T J\nAND D J\nRUN\n
			//

			command[command[i+1]] = int(str[not_d_j_walk])
			fmt.Println(command[command[i+1]])
			not_d_j_walk++

			i += 2
		case 4: //output
			// if value1 == 10 {
			// 	output += "\n"
			// }
			output += string(value1)

			// fmt.Println(numStr, " -> ", value1)
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(command[i+1]))
			fmt.Println(string(value1), value1, " output ", i, opcode)
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

	fmt.Println(output, []byte(output))
	str := ""
	for _, byte := range []byte(output) {
		str += strconv.Itoa(int(byte))
	}
	println(str)
}

func main() {
	part1()
}
