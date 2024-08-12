package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type amplifier struct {
	field           []int
	inputsLeft      []int
	curCommandIndex int
	lastOutput      int
	relativeBase    int
}

func main() {
	part1()
}

func part1() (int, error) {
	base := 0

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

func permutations() [120][5]int {

	index := 0
	permutations := [120][5]int{}
	curChoice := [5]int{}
	for i, p := range [5]int{5, 6, 7, 8, 9} {
		curChoice[0] = p
		remainingNumbers := []int{5, 6, 7, 8, 9}
		// remainingNumbers = slices.Delete(remainingNumbers, i, i+1)
		remainingNumbers = append(remainingNumbers[:i], remainingNumbers[i+1:]...)
		// fmt.Println(remainingNumbers)
		for j, k := range remainingNumbers {
			curChoice[1] = k
			remainingNumbers2 := []int{remainingNumbers[0], remainingNumbers[1], remainingNumbers[2], remainingNumbers[3]}
			// remainingNumbers = slices.Delete(remainingNumbers, j, j+1)
			remainingNumbers2 = append(remainingNumbers2[:j], remainingNumbers2[j+1:]...)
			// fmt.Println(remainingNumbers2)
			for l, m := range remainingNumbers2 {
				curChoice[2] = m
				remainingNumbers3 := []int{remainingNumbers2[0], remainingNumbers2[1], remainingNumbers2[2]}
				// remainingNumbers = slices.Delete(remainingNumbers, l, l+1)
				remainingNumbers3 = append(remainingNumbers3[:l], remainingNumbers3[l+1:]...)
				// fmt.Println(remainingNumbers3)
				for n, o := range remainingNumbers3 {
					// fmt.Println(n, n+1)
					curChoice[3] = o
					q := remainingNumbers3[(n+1)%len(remainingNumbers3)]
					curChoice[4] = q
					permutations[index] = [5]int{curChoice[0], curChoice[1], curChoice[2], curChoice[3], curChoice[4]}
					index++
				}
			}
		}
	}

	return permutations
}
