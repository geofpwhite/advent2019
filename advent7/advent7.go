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
}

func part2() {
	amplifiers := make([]*amplifier, 5)
	permutations := permutations()
	maxOutput := 0
	maxphase := 0
	for k, phaseSettings := range permutations {
		for i := 0; i < 5; i++ {
			amplifiers[i] = &amplifier{}
			amplifiers[i].field, _ = parse()
			amplifiers[i].inputsLeft = append(amplifiers[i].inputsLeft, phaseSettings[i])
		}
		amplifiers[0].inputsLeft = []int{amplifiers[0].inputsLeft[0], 0}
		curAmpIndex := 0
		lastOutput := 0
		// fmt.Println(phaseSettings)
		for {
			empties := 0
			for _, amp := range amplifiers {
				if len(amp.inputsLeft) == 0 && amp.field[amp.curCommandIndex] == 3 {
					empties++
				}
			}
			if empties == 5 {
				println("empty")
				break
			}

			curAmplifier := amplifiers[curAmpIndex]
			command := curAmplifier.field[curAmplifier.curCommandIndex]
			numStr := strconv.Itoa(command)
			for len(numStr) < 5 {
				numStr = "0" + numStr
			}
			opcode, _ := strconv.Atoi(numStr[3:])
			// println(opcode)
			paramModes := []byte{numStr[2], numStr[1], numStr[0]}
			var value1, value2 int
			if opcode == 99 && curAmpIndex != 4 {
				fmt.Println(curAmpIndex)
				curAmpIndex = (curAmpIndex + 1) % 5
				// amplifiers[curAmpIndex].inputsLeft = append(amplifiers[curAmpIndex].inputsLeft, lastOutput)
				fmt.Println(curAmpIndex)
				continue
			} else if opcode == 99 && curAmpIndex == 4 {
				// println(lastOutput)
				if curAmplifier.lastOutput > maxOutput {
					maxOutput = curAmplifier.lastOutput
					maxphase = k
				}
				break
			}

			if paramModes[0] == '1' {
				value1 = curAmplifier.field[curAmplifier.curCommandIndex+1]
			} else {
				// fmt.Println(command, opcode, curAmpIndex)
				value1 = curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+1]]
			}
			if paramModes[1] == '1' {
				value2 = curAmplifier.field[curAmplifier.curCommandIndex+2]
			} else if opcode != 3 && opcode != 4 {
				value2 = curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+2]]
			}
			switch opcode {
			case 1: //addition
				curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+3]] = value1 + value2
				curAmplifier.curCommandIndex += 4
			case 2: //multiplication
				curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+3]] = value1 * value2
				curAmplifier.curCommandIndex += 4
			case 3: //input
				// fmt.Println(curAmplifier.inputsLeft)
				if len(curAmplifier.inputsLeft) == 0 {
					curAmpIndex = (curAmpIndex + 1) % 5
					// println("continued")
					// fmt.Println(lastOutput)
					continue
				}
				curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+1]] = curAmplifier.inputsLeft[0]
				curAmplifier.inputsLeft = curAmplifier.inputsLeft[1:]
				curAmplifier.curCommandIndex += 2
				curAmpIndex = (curAmpIndex + 1) % 5
			case 4: //output
				// outputChannel <- strconv.Itoa(curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+1]])
				amplifiers[(curAmpIndex+1)%5].inputsLeft = append(amplifiers[(curAmpIndex+1)%5].inputsLeft, curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+1]])
				fmt.Println(curAmpIndex, curAmplifier.curCommandIndex)
				fmt.Println("addInput: " + strconv.Itoa(curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+1]]))
				lastOutput = curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+1]]
				curAmplifier.lastOutput = lastOutput
				curAmplifier.curCommandIndex += 2
				curAmpIndex = (curAmpIndex + 1) % 5
				// curAmpIndex = (curAmpIndex + 1) % 5
			case 5:
				if value1 != 0 {
					curAmplifier.curCommandIndex = value2
				} else {
					curAmplifier.curCommandIndex += 3
				}
			case 6:
				if value1 == 0 {
					curAmplifier.curCommandIndex = value2
				} else {
					curAmplifier.curCommandIndex += 3
				}
			case 7:
				if value1 < value2 {
					curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+3]] = 1
				} else {
					curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+3]] = 0
				}
				curAmplifier.curCommandIndex += 4
			case 8:
				if value1 == value2 {
					curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+3]] = 1
				} else {
					curAmplifier.field[curAmplifier.field[curAmplifier.curCommandIndex+3]] = 0
				}
				curAmplifier.curCommandIndex += 4
			case 99:
				fmt.Println("output " + strconv.Itoa(lastOutput))
				break
			}
		}
		// outputNum, err := strconv.Atoi(output)
		// if err != nil {
		// panic("problem, output is " + output)
		// return -1, fmt.Errorf("bad output somehow")
		// }
		// return outputNum, nil

	}
	fmt.Println(maxOutput, permutations[maxphase])
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
		fmt.Println(remainingNumbers)
		for j, k := range remainingNumbers {
			curChoice[1] = k
			remainingNumbers2 := []int{remainingNumbers[0], remainingNumbers[1], remainingNumbers[2], remainingNumbers[3]}
			// remainingNumbers = slices.Delete(remainingNumbers, j, j+1)
			remainingNumbers2 = append(remainingNumbers2[:j], remainingNumbers2[j+1:]...)
			fmt.Println(remainingNumbers2)
			for l, m := range remainingNumbers2 {
				curChoice[2] = m
				remainingNumbers3 := []int{remainingNumbers2[0], remainingNumbers2[1], remainingNumbers2[2]}
				// remainingNumbers = slices.Delete(remainingNumbers, l, l+1)
				remainingNumbers3 = append(remainingNumbers3[:l], remainingNumbers3[l+1:]...)
				fmt.Println(remainingNumbers3)
				for n, o := range remainingNumbers3 {
					fmt.Println(n, n+1)
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

func main() {
	part2()
}

func part1() {

	index := 0
	permutations := [120][5]int{}
	curChoice := [5]int{}
	for i := range [5]int{} {
		curChoice[0] = i
		remainingNumbers := []int{0, 1, 2, 3, 4}
		// remainingNumbers = slices.Delete(remainingNumbers, i, i+1)
		remainingNumbers = append(remainingNumbers[:i], remainingNumbers[i+1:]...)
		fmt.Println(remainingNumbers)
		for j, k := range remainingNumbers {
			curChoice[1] = k
			remainingNumbers2 := []int{remainingNumbers[0], remainingNumbers[1], remainingNumbers[2], remainingNumbers[3]}
			// remainingNumbers = slices.Delete(remainingNumbers, j, j+1)
			remainingNumbers2 = append(remainingNumbers2[:j], remainingNumbers2[j+1:]...)
			fmt.Println(remainingNumbers2)
			for l, m := range remainingNumbers2 {
				curChoice[2] = m
				remainingNumbers3 := []int{remainingNumbers2[0], remainingNumbers2[1], remainingNumbers2[2]}
				// remainingNumbers = slices.Delete(remainingNumbers, l, l+1)
				remainingNumbers3 = append(remainingNumbers3[:l], remainingNumbers3[l+1:]...)
				fmt.Println(remainingNumbers3)
				for n, o := range remainingNumbers3 {
					fmt.Println(n, n+1)
					curChoice[3] = o
					q := remainingNumbers3[(n+1)%len(remainingNumbers3)]
					curChoice[4] = q
					permutations[index] = [5]int{curChoice[0], curChoice[1], curChoice[2], curChoice[3], curChoice[4]}
					index++
				}
			}
		}
	}
	maximum := 0
	prev := 0
	maxPhase := [5]int{}
	for _, phaseSettings := range permutations {
		var err error
		prev, err = advent7(phaseSettings[0], 0)
		if err != nil {
			fmt.Println(err)
			continue
		}
		prev, err = advent7(phaseSettings[1], prev)
		if err != nil {
			fmt.Println(err)
			continue
		}
		prev, err = advent7(phaseSettings[2], prev)
		if err != nil {
			fmt.Println(err)
			continue
		}
		prev, err = advent7(phaseSettings[3], prev)
		if err != nil {
			fmt.Println(err)
			continue
		}
		prev, err = advent7(phaseSettings[4], prev)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if prev > maximum {
			maximum = prev
			maxPhase = phaseSettings
		}
	}
	fmt.Println(maximum, maxPhase)
}

func part2thread() {
	command, length := parse()
	// fmt.Println(command)
	output := ""
	for i := 0; i < length; {
		numStr := strconv.Itoa(command[i])
		for len(numStr) < 5 {
			numStr = "0" + numStr
		}
		opcode, _ := strconv.Atoi(numStr[3:])
		paramModes := []byte{numStr[2], numStr[1], numStr[0]}
		var value1, value2 int
		if opcode == 99 {
			break
		}

		if paramModes[0] == '1' {
			value1 = command[i+1]
		} else { //if opcode != 4 && opcode != 3 {
			fmt.Println(command, opcode)
			value1 = command[command[i+1]]
		}
		if paramModes[1] == '1' {
			value2 = command[i+2]
		} else if opcode != 3 && opcode != 4 {
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
			// command[command[i+1]], _ = strconv.Atoi(<-inputChannel)
			i += 2
		case 4: //output
			output += strconv.Itoa(command[command[i+1]])
			// outputChannel <- strconv.Itoa(command[command[i+1]])
			fmt.Println(output)
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
	}
	// outputNum, err := strconv.Atoi(output)
	// if err != nil {
	// panic("problem, output is " + output)
	// return -1, fmt.Errorf("bad output somehow")
	// }
	// return outputNum, nil

}

func advent7(firstInput, secondInput int) (int, error) {
	command, length := parse()
	// fmt.Println(command)
	output := ""
	for i := 0; i < length; {
		numStr := strconv.Itoa(command[i])
		for len(numStr) < 5 {
			numStr = "0" + numStr
		}
		opcode, _ := strconv.Atoi(numStr[3:])
		paramModes := []byte{numStr[2], numStr[1], numStr[0]}
		var value1, value2 int
		if opcode == 99 {
			break
		}

		if paramModes[0] == '1' {
			value1 = command[i+1]
		} else { //if opcode != 4 && opcode != 3 {
			fmt.Println(command, opcode)
			value1 = command[command[i+1]]
		}
		if paramModes[1] == '1' {
			value2 = command[i+2]
		} else if opcode != 3 && opcode != 4 {
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
			if firstInput == -1 {
				command[command[i+1]] = secondInput
			} else {
				command[command[i+1]] = firstInput
				firstInput = -1
			}
			i += 2
		case 4: //output
			output += strconv.Itoa(command[command[i+1]])
			fmt.Println(output)
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
	}
	outputNum, err := strconv.Atoi(output)
	if err != nil {
		// panic("problem, output is " + output)
		return -1, fmt.Errorf("bad output somehow")
	}
	return outputNum, nil

}
