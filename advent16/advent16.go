package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// starting at step 0
func patternByElement(element int) []int {
	og := [4]int{0, 1, 0, -1}
	pattern := []int{}
	for i := 0; i < 4; i++ {
		for j := 0; j < element+1; j++ {
			pattern = append(pattern, og[i])
		}
	}
	pattern = append(pattern[1:], pattern[0])
	return pattern
}

func main() {
	part2()
}

func parse() []string {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
	lines := strings.Split(string(content), "\n")
	for i := range lines {
		lines[i] = strings.ReplaceAll(lines[i], "\n", "")
		fmt.Println(lines[i])
	}
	return lines
}

func part1() {
	line := parse()[0]

	newLine := ""
	for phase := 0; phase < 100; phase++ {
		for i := range line {
			pattern := patternByElement(i)
			sum := 0
			for j, char2 := range line {
				num, _ := strconv.Atoi(string(char2))
				num *= pattern[j%len(pattern)]
				sum += num
			}
			sumStr := strconv.Itoa(sum)
			newLine += sumStr[len(sumStr)-1:]

		}
		line = newLine
		newLine = ""
	}
	fmt.Println(line[:8])
}

func part2() {
	line := parse()[0]
	x := line
	// patterns := [10000][]int{}
	for range 9999 {
		line += x
		// patterns[i] = patternByElement(i)
	}
	// patterns[9999] = patternByElement(9999)
	offset, _ := strconv.Atoi(line[:7])
	line = line[offset:]
	signal := make([]int, len(line))
	for i, char := range line {
		signal[i], _ = strconv.Atoi(string(char))
	}
	// newLine := ""
	//because phase is in second half of the signal, the patttern we use will be the first half of 0's and the second half of 1's
	//had to look this up but the logic is very cool

	for phase := 0; phase < 100; phase++ {
		cumulative := 0
		println(phase)
		fmt.Println(signal[:8])
		newSignal := make([]int, len(line))
		for i := len(signal) - 1; i > -1; i-- {
			// num, _ := strconv.Atoi(string(line[i]))
			cumulative = (cumulative + signal[i]) % 10
			newSignal[i] = int(math.Abs(float64(cumulative)))
			// line = line[:i] + strconv.Itoa(cumulative) + line[i+1:]
		}
		signal = newSignal
		// line = newLine
		// newLine = ""
	}
	fmt.Println(signal[:8])
}
