package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	adventhelperfunctions "github.com/geofpwhite/adventHelperFunctions"
	//	adventhelperfunctions "github.com/geofpwhite/adventHelperFunctions"
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

type direction int

const NORTH direction = 1
const SOUTH direction = 2
const WEST direction = 3
const EAST direction = 4

type droid struct {
	x, y, xTry, yTry   int
	field              map[[2]int]int
	directionsTried    map[[2]int][]direction
	curDirectionTrying direction
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Sqrt((math.Pow(math.Abs(float64(x2-x1)), 2)) + (math.Pow(math.Abs(float64(y2-y1)), 2)))
}

func (r *droid) moveForward(dir direction) {
	switch dir {
	case NORTH:
		r.xTry, r.yTry = r.x, r.y
		r.yTry++
	case SOUTH:
		r.xTry, r.yTry = r.x, r.y
		r.yTry--
	case EAST:
		r.xTry, r.yTry = r.x, r.y
		r.xTry++
	case WEST:
		r.xTry, r.yTry = r.x, r.y
		r.xTry--
	}
}

func (r *droid) getNeighbors(x, y int) [][2]int {
	neighbors := [][2]int{
		{x, y + 1},
		{x, y - 1},
		{x + 1, y},
		{x - 1, y},
	}
	for i := 3; i > -1; i-- {
		if r.field[neighbors[i]] == -1 {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
		}
	}
	return neighbors

}
func (r *droid) bestDirection() direction {
	// dirTried := r.directionsTried[[2]int{r.x, r.y}]
	val := []bool{true, true, true, true}
	dirs := []direction{NORTH, SOUTH, EAST, WEST}
	neighbors := [][2]int{
		{r.x, r.y + 1},
		{r.x, r.y - 1},
		{r.x + 1, r.y},
		{r.x - 1, r.y},
	}
	for i := range neighbors {
		if r.field[neighbors[i]] == -1 {
			val[i] = false
		}
	}
	for i := 3; i > -1; i-- {
		if !val[i] {
			neighbors = append(neighbors[:i], neighbors[i+1:]...)
			dirs = append(dirs[:i], dirs[i+1:]...)
			val = append(val[:i], val[i+1:]...)
		}
	}
	if len(neighbors) == 1 {
		return dirs[0]
	}
	//if we get to this line it means that the length of dirs/val/neighbors is > 1
	// for i := range dirs {
	// 	if r.field[neighbors[i]] == 1 {
	// 		val[i] = false
	// 	}
	// }

	// if slices.Contains(val, true) {
	// 	for i := len(val) - 1; i > -1; i-- {
	// 		if !val[i] {
	// 			neighbors = append(neighbors[:i], neighbors[i+1:]...)
	// 			dirs = append(dirs[:i], dirs[i+1:]...)
	// 			val = append(val[:i], val[i+1:]...)
	// 		}
	// 	}

	// 	return dirs[rand.Intn(len(dirs))]
	// } else {
	return dirs[rand.Intn(len(dirs))]
	// return
	// }

}

type node struct {
	x, y      int
	neighbors map[[2]int]int
}
type graph map[[2]int]*node

func part1() {
	fmt.Println("f")
	base := 0
	command, length := parse()
	rob := droid{field: map[[2]int]int{{0, 0}: 1}, directionsTried: make(map[[2]int][]direction)}
	unvisited := make(map[[2]int]bool)
	visited := make(map[[2]int]bool)
	command = append(command, make([]int, 100)...)
	check := false
	// cycles := [4]int{1, 4, 2, 3}
	// cycle := 1
	// fmt.Println(command)
	defer func() {
		x := adventhelperfunctions.MapToImage[int](rob.field)
		filestr := ""
		for _, line := range x {
			for _, char := range line {
				char2 := strconv.Itoa(char)
				if char2 == "-1" {
					char2 = "#"
				}
				if char2 == "1" {
					char2 = " "
				}
				filestr += (char2)
			}
			filestr += "\n"
		}
		file, _ := os.Create("output.txt")
		file.WriteString(filestr)
		file.Close()
		g := make(graph)
		dist := make(map[[2]int]int)

		for coords, value := range rob.field {
			if value < 1 {
				continue
			}
			dist[coords] = -1

			if g[coords] == nil {
				g[coords] = &node{coords[0], coords[1], make(map[[2]int]int)}
			}
			neighbors := rob.getNeighbors(coords[0], coords[1])
			for _, coords2 := range neighbors {
				g[coords].neighbors[coords2] = 1
			}
		}
		start := g[[2]int{-16, -12}]
		dist[[2]int{-16, -12}] = 0
		queue := []node{*start}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			// if rob.field[[2]int{cur.x, cur.y}] == 2 {
			// 	fmt.Println(dist[[2]int{cur.x, cur.y}])
			// 	return
			// }
			for coords := range cur.neighbors {
				if dist[coords] == -1 || dist[[2]int{cur.x, cur.y}]+1 < dist[coords] {
					queue = append(queue, *g[coords])
					dist[coords] = dist[[2]int{cur.x, cur.y}] + 1
				}
			}
			// fmt.Println(dist[two])

		}

		x = adventhelperfunctions.MapToImage[int](dist)
		filestr = ""
		for _, line := range x {
			for _, char := range line {
				char2 := strconv.Itoa(char)
				if char2 == "-1" || char == 0 {
					char2 = "######"
				}
				filestr += (char2)
				for i := len(char2); i < 6; i++ {
					filestr += " "
				}
			}
			filestr += "\n"
		}
		file2, _ := os.Create("output2.txt")
		file2.WriteString(filestr)
		file2.Close()
		max := 0
		var co [2]int
		for coo, val := range dist {
			if val > max {
				max = val
				co = coo
			}
		}
		fmt.Println(max, co, rob.field[co])
	}()
	x := 0
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
		} else { //if opcode != 3 {
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

			rob.curDirectionTrying = rob.bestDirection()
			// fmt.Println(rob.curDirectionTrying)
			rob.moveForward(rob.curDirectionTrying)
			// for j := rand.Intn(4); ; j = (j + 1) % 4 {

			// 	if rob.field[getCoordsForDirection(direction(cycles[j]), rob.x, rob.y)] != -1 &&
			// 		(!slices.Contains(rob.directionsTried[[2]int{rob.x, rob.y}], direction(cycles[j])) || (len(rob.directionsTried[[2]int{rob.x, rob.y}]) >= 4 || rob.allNeighborsVisited(rob.x, rob.y))) {

			// 		rob.curDirectionTrying = direction(cycles[j])
			// 		fmt.Println(direction(cycles[j]), rob.x, rob.y)
			// 		rob.moveForward(rob.curDirectionTrying)
			// 		break
			// 	}
			// }
			// cycle = (cycle + 1) % 4
			if paramModes[0] == '2' {
				command[base+command[i+1]] = int(rob.curDirectionTrying)
			} else {
				command[command[i+1]] = int(rob.curDirectionTrying)
			}

			i += 2
		case 4: //output

			// if rob.x > 100 || rob.y > 100 ||
			// 	rob.x < -100 || rob.y < -100 {
			// 	println("break")
			// 	return
			// }
			// bre := false
			x++
			if check {
				if len(unvisited) == 0 {
					return
				}

			}
			delete(unvisited, [2]int{rob.xTry, rob.yTry})
			visited[[2]int{rob.xTry, rob.yTry}] = true
			if value1 == 1 {
				rob.field[[2]int{rob.xTry, rob.yTry}] = 1
				rob.directionsTried[[2]int{rob.x, rob.y}] = append(rob.directionsTried[[2]int{rob.x, rob.y}], rob.curDirectionTrying)
				n := rob.getNeighbors(rob.xTry, rob.yTry)
				for _, c := range n {
					if !visited[[2]int{c[0], c[1]}] {
						unvisited[[2]int{c[0], c[1]}] = true
					}
				}

				rob.x, rob.y = rob.xTry, rob.yTry
			} else if value1 == 2 {
				rob.field[[2]int{rob.xTry, rob.yTry}] = 2
				rob.directionsTried[[2]int{rob.x, rob.y}] = append(rob.directionsTried[[2]int{rob.x, rob.y}], rob.curDirectionTrying)
				n := rob.getNeighbors(rob.xTry, rob.yTry)
				for _, c := range n {
					if rob.field[[2]int{c[0], c[1]}] == 0 {
						unvisited[[2]int{c[0], c[1]}] = true
					}
				}
				rob.x, rob.y = rob.xTry, rob.yTry
				check = true
				// return

			} else {
				rob.field[[2]int{rob.xTry, rob.yTry}] = -1
				rob.directionsTried[[2]int{rob.x, rob.y}] = append(rob.directionsTried[[2]int{rob.x, rob.y}], rob.curDirectionTrying)
			}
			// fmt.Println(value1, rob.curDirectionTrying)
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
			fmt.Printf("%d\n", base)
			base += value1
			// base += command[i+1]
			// fmt.Printf("%d\n", base)
			// fmt.Println(strconv.Itoa(base) + " + " + strconv.Itoa(value1))
			i += 2
		case 99:
			break
		}

	}
	fmt.Println(rob.field)

}
