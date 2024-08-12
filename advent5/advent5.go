package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// advent2()
	advent5()
}

func advent2() {
	content, _ := os.ReadFile("inputadvent2.txt")
	lines := strings.Split(string(content), "\n")
	command := []int{}
	for _, numStr := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("ahhhh")
		}
		command = append(command, num)
	}
	length := len(command)
	fmt.Println(command)
	command[1], command[2] = 12, 2
	for i := 0; i < length; {
		switch command[i] {
		case 1:
			command[command[i+3]] = command[command[i+1]] + command[command[i+2]]
			i += 4
		case 2:
			command[command[i+3]] = command[command[i+1]] * command[command[i+2]]
			i += 4
		case 3:

		case 99:
			break
		}
	}
	fmt.Println(command[0])

}

func advent5() {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	command := []int{}
	for _, numStr := range strings.Split(lines[0], ",") {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			panic("ahhhh")
		}
		command = append(command, num)
	}
	length := len(command)
	fmt.Println(command)
	output := ""
	for i := 0; i < length; {
		numStr := strconv.Itoa(command[i])
		for len(numStr) < 5 {
			numStr = "0" + numStr
		}
		opcode, _ := strconv.Atoi(numStr[3:])
		paramModes := []byte{numStr[2], numStr[1], numStr[0]}
		var value1, value2 int

		if paramModes[0] == '1' {
			value1 = command[i+1]
		} else {
			value1 = command[command[i+1]]
		}
		if opcode == 99 {
			break
		}
		if paramModes[1] == '1' {
			value2 = command[i+2]
		} else {
			value2 = command[command[i+2]]
		}
		switch opcode {
		case 1: //addition
			command[command[i+3]] = value1 + value2
			i += 4
		case 2: //multiplication
			command[command[i+3]] = value1 * value2
			i += 4
		case 3: //input
			command[command[i+1]] = 5
			i += 2
		case 4: //output
			output += strconv.Itoa(command[command[i+1]])
			fmt.Println(output, strconv.Itoa(command[command[i+1]]))
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
				command[command[i+3]] = 1
			} else {
				command[command[i+3]] = 0
			}
			i += 4
		case 8:

			if value1 == value2 {
				command[command[i+3]] = 1
			} else {
				command[command[i+3]] = 0
			}
			i += 4
		case 99:
			break
		}
		fmt.Println(i, length, command[i:i+4])
	}
	fmt.Println(output)

}
