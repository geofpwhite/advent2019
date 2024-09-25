package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type coords [2]int

type layer struct {
	outer *layer
	inner *layer
	bugs  map[coords]bool
}
type layers map[[3]int]bool

func adjacentAbove(layerIndex, x, y int) [][3]int {
	switch x {
	case 0:
		return [][3]int{
			{layerIndex + 1, 1, 2},
		}
	case 3:
		if y == 2 {
			return [][3]int{
				{layerIndex - 1, 4, 0},
				{layerIndex - 1, 4, 1},
				{layerIndex - 1, 4, 2},
				{layerIndex - 1, 4, 3},
				{layerIndex - 1, 4, 4},
			}
		} else {
			return [][3]int{{layerIndex, x - 1, y}}
		}
	default:
		return [][3]int{{layerIndex, x - 1, y}}
	}
}

func adjacentBelow(layerIndex, x, y int) [][3]int {
	switch x {
	case 4:
		return [][3]int{
			{layerIndex + 1, 3, 2},
		}
	case 1:
		if y == 2 {
			return [][3]int{
				{layerIndex - 1, 0, 0},
				{layerIndex - 1, 0, 1},
				{layerIndex - 1, 0, 2},
				{layerIndex - 1, 0, 3},
				{layerIndex - 1, 0, 4},
			}
		} else {
			return [][3]int{{layerIndex, x + 1, y}}
		}
	default:
		return [][3]int{{layerIndex, x + 1, y}}
	}

}
func adjacentLeft(layerIndex, x, y int) [][3]int {
	switch y {
	case 0:
		return [][3]int{
			{layerIndex + 1, 2, 1},
		}
	case 3:
		if x == 2 {
			return [][3]int{
				{layerIndex - 1, 0, 4},
				{layerIndex - 1, 1, 4},
				{layerIndex - 1, 2, 4},
				{layerIndex - 1, 3, 4},
				{layerIndex - 1, 4, 4},
			}
		} else {
			return [][3]int{{layerIndex, x, y - 1}}
		}
	default:
		return [][3]int{{layerIndex, x, y - 1}}
	}

}
func adjacentRight(layerIndex, x, y int) [][3]int {
	switch y {
	case 4:
		return [][3]int{
			{layerIndex + 1, 2, 3},
		}
	case 1:
		if x == 2 {
			return [][3]int{
				{layerIndex - 1, 0, 0},
				{layerIndex - 1, 1, 0},
				{layerIndex - 1, 2, 0},
				{layerIndex - 1, 3, 0},
				{layerIndex - 1, 4, 0},
			}
		} else {
			return [][3]int{{layerIndex, x, y + 1}}
		}
	default:
		return [][3]int{{layerIndex, x, y + 1}}
	}

}

func (ls layers) neighborsToCheck(layerIndex, x, y int) [][3]int {
	neighborsToCheck := [][3]int{}
	neighborsToCheck = append(neighborsToCheck, adjacentAbove(layerIndex, x, y)...)
	neighborsToCheck = append(neighborsToCheck, adjacentBelow(layerIndex, x, y)...)
	neighborsToCheck = append(neighborsToCheck, adjacentLeft(layerIndex, x, y)...)
	neighborsToCheck = append(neighborsToCheck, adjacentRight(layerIndex, x, y)...)
	return neighborsToCheck
	// switch [2]int{x, y} {
	// case [2]int{0, 0}, [2]int{4, 4}, [2]int{0, 4}, [2]int{4, 0},[2]int{}: //4 adjacent
	// case [2]int{1, 1}, [2]int{1, 3}, [2]int{3, 3}, [2]int{3, 1}:
	// case [2]int{0, 0}, [2]int{4, 4}, [2]int{0, 4}, [2]int{4, 0}:
	// }
}
func (ls layers) morph() layers {
	var emptyButBugNeighbors map[[3]int]int = map[[3]int]int{}
	newLayers := make(layers)

	for coords := range ls {
		bugNeighbors := 0
		neighbors := ls.neighborsToCheck(coords[0], coords[1], coords[2])
		for _, coord := range neighbors {
			if ls[coord] {
				bugNeighbors++
			} else {
				emptyButBugNeighbors[coord]++
			}
		}
		if bugNeighbors == 1 {
			newLayers[coords] = true
		}
	}
	for coords, bugNeighbors := range emptyButBugNeighbors {
		if bugNeighbors == 1 || bugNeighbors == 2 {
			newLayers[coords] = true
		}
	}
	return newLayers
}
func part2() {
	ls := layers{}
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if len(line) > 0 {
			for j, char := range line {
				if char == '#' {
					ls[[3]int{0, i, j}] = true
				}
			}
		}
	}
	for i := 0; i < 200; i++ {
		fmt.Println(len(ls))
		ls = ls.morph()
	}
	fmt.Println(len(ls))
	fmt.Println(ls.neighborsToCheck(0, 0, 0))
}

type grid []string

func main() {
	part2()
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
