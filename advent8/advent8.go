package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse() {

}

func main() {
	content, _ := os.ReadFile("input.txt")
	line := strings.Replace(string(content), "\r", "", -1)
	zeroes := []int{}
	layers := [][25][6]rune{}
	ones, twos := []int{}, []int{}
	curZeroes, curOnes, curTwos := 0, 0, 0
	curLayer := [25][6]rune{}
	index1, index2 := 0, 0
	for i, char := range line {
		if i%150 == 0 && i != 0 {
			zeroes = append(zeroes, curZeroes)
			ones = append(ones, curOnes)
			twos = append(twos, curTwos)
			curZeroes = 0
			curOnes = 0
			curTwos = 0
			index1 = 0
			index2 = 0
			layers = append(layers, curLayer)
			curLayer = [25][6]rune{}
		}
		curLayer[index1][index2] = char
		if index1 == 24 {
			index2++
			index1 = 0
		} else {
			index1++
		}
		if char == '0' {
			curZeroes++
		}
		if char == '1' {
			curOnes++
		}
		if char == '2' {
			curTwos++
		}
	}
	minimum := 10000000
	index := 0
	for i, num := range zeroes {
		if num < minimum && num != 0 {
			index = i
			minimum = num
		}
	}
	fmt.Println(ones[index] * twos[index])
	final := [25][6]int{}
	for i := 0; i < 25; i++ {
		for j := 0; j < 6; j++ {
			index := 0
			for layers[index][i][j] == '2' {
				index++
				println(i, j, index)
			}
			final[i][j], _ = strconv.Atoi(string(layers[index][i][j]))
		}
	}
	fmt.Println(final)
	printString := ""
	for j := 0; j < 6; j++ {
		for i := 0; i < 25; i++ {
			printString += strconv.Itoa(final[i][j])
		}
		fmt.Println(strings.Replace(printString, "0", " ", -1))
		printString = ""
	}
}
