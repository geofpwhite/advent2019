package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type coords [2]int
type threeDcoords [3]int
type grid []string

type bugCube map[threeDcoords]bool

func main() {
	part1()
}

func (g grid) toString() string {
	return strings.Join(g, "")
}

func (g grid) morph() grid {
	g2 := make(grid, len(g))
	for i, line := range g {
		for j, char := range line {
			neighbors := g.neighbors(i, j)
			if char == '#' {
				if neighbors != 1 {
					g2[i] += "."
				} else {
					g2[i] += "#"
				}
			} else {
				if neighbors == 1 || neighbors == 2 {
					g2[i] += "#"
				} else {
					g2[i] += "."
				}

			}
		}
	}
	return g2
}

func (g grid) neighbors(x, y int) int {
	neighbors := []coords{
		{x + 1, y},
		{x, y + 1},
		{x - 1, y},
		{x, y - 1},
	}
	sum := 0
	for _, c := range neighbors {
		if c[0] > -1 && c[1] > -1 && c[0] < len(g) && c[1] < len(g[c[0]]) && g[c[0]][c[1]] == '#' {
			sum++
		}
	}
	return sum
}

func part1() {
	g := parse()
	prevLayouts := map[string]bool{}
	for !prevLayouts[g.toString()] {
		prevLayouts[g.toString()] = true
		g = g.morph()
	}
	fmt.Println(g.toString())
	sum := 0
	for i, char := range g.toString() {
		pow := int(math.Pow(2, float64(i)))
		if char == '#' {
			sum += pow
		}
	}
	fmt.Println(sum)
}

func parse() grid {
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
