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

func part1() {
	funcs, params, ary := parse1()

	og := make([]int, len(ary))
	check := make([]int, 1000000)
	for i := range check {
		check[i] = i + 10007
	}
	ary = append(ary, check...)
	copy(og, ary)
	for i := range funcs {
		ary = funcs[i](ary, params[i])
		fmt.Println(funcs[i], params[i])
		fmt.Println(ary[2019])
	}
	// fmt.Println(119315717514046 - slices.Index(ary, 2020))
	longestStreak, curStreak := 0, 1

	for i, num := range ary[:10000000] {
		fmt.Println(i, num)
		if i > 0 && num == ary[i-1]-1 {
			curStreak++
			longestStreak = max(longestStreak, curStreak)
		} else {
			curStreak = 1
		}
	}
	fmt.Println(longestStreak)

	// fmt.Println(ary[119315717514046-2020])
}

func newStack(ary []int, in int) []int {
	length := len(ary)
	n := make([]int, length)

	for i := 0; i < length; i++ {
		n[i] = ary[length-1-i]
	}
	return n
}

func cut(ary []int, index int) []int {
	var length int = len(ary)
	n := make([]int, length)
	if index < 0 {
		index = length + index
	}
	copy(n[:index], ary[length-index:])
	copy(n[index:], ary[:length-index])
	return n
}
func deal(ary []int, increment int) []int {
	length := len(ary)
	l2 := length - 1
	n := make([]int, length)
	index := 0
	for l2 > 0 {
		n[length-1-index] = ary[l2]
		index = (index + increment) % length
		l2--
	}
	return n
}

func parse1() ([]func(ary []int, in int) []int, []int, []int) {
	funcs := []func(ary []int, in int) []int{}
	params := []int{}
	ary := make([]int, 10007)
	for i := range ary {
		ary[10006-i] = i
	}
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "new stack") {
			funcs = append(funcs, newStack)
			params = append(params, -1)
		} else if strings.Contains(line, "cut") {
			funcs = append(funcs, cut)
			num, _ := strconv.Atoi(line[4:])
			params = append(params, num)
		} else {
			funcs = append(funcs, deal)
			num, _ := strconv.Atoi(line[20:])
			params = append(params, num)
		}
	}
	return funcs, params, ary
}

func parse2() ([]func(ary []int, in int) []int, []int, []int) {
	funcs := []func(ary []int, in int) []int{}
	params := []int{}
	ary := make([]int, 119315717514047)
	for i := range ary {
		ary[119315717514046-i] = i
	}
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "new stack") {
			funcs = append(funcs, newStack)
			params = append(params, -1)
		} else if strings.Contains(line, "cut") {
			funcs = append(funcs, cut)
			num, _ := strconv.Atoi(line[4:])
			params = append(params, num)
		} else {
			funcs = append(funcs, deal)
			num, _ := strconv.Atoi(line[20:])
			params = append(params, num)
		}
	}
	return funcs, params, ary
}
