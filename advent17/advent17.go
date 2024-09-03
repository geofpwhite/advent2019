package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
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

type field []string

type node struct {
	up, down, left, right      *node
	uEdge, dEdge, lEdge, rEdge int
	x, y                       int
}

type graph map[[2]int]*node

var lines field

type pathAndPosition struct {
	path      *pathAndPosition //previous node that was visited
	position  [2]int
	direction rune // either '<','>','v', or '^'
	edge      int
}

func (p *pathAndPosition) toString() (str string) {
	str += strconv.Itoa(p.position[0]) + " " + strconv.Itoa(p.position[1])
	visited := ""
	visitedNum := 0
	for p.path != nil {
		// prev := p.path
		visited = "( " + strconv.Itoa(p.position[0]) + " , " + strconv.Itoa(p.position[1]) + " ) " + visited
		visitedNum++
		prevDir, dir := p.path.direction, p.direction
		directionToAdd := ""
		if prevDir == dir {
			for prevDir == dir {
				p.edge += p.path.edge
				p.path = p.path.path
				prevDir, dir = p.path.direction, p.direction
			}
		}
		switch prevDir {
		case '>':
			if dir == '^' {
				directionToAdd = "L"
			} else if dir == 'v' {
				directionToAdd = "R"
			} else if dir == '<' {
				directionToAdd = "RR"
			}
		case '<':
			if dir == '^' {
				directionToAdd = "R"
			} else if dir == 'v' {
				directionToAdd = "L"
			} else if dir == '>' {
				directionToAdd = "RR"
			}
		case 'v':
			if dir == '<' {
				directionToAdd = "R"
			} else if dir == '>' {
				directionToAdd = "L"
			} else if dir == '^' {
				directionToAdd = "RR"
			}
		case '^':
			if dir == '<' {
				directionToAdd = "L"
			} else if dir == '>' {
				directionToAdd = "R"
			} else if dir == 'v' {
				directionToAdd = "RR"
			}
		}
		str = directionToAdd + strconv.Itoa(p.edge) + " " + str
		p = p.path
	}
	str += "\n" + visited + strconv.Itoa(visitedNum)
	return
}

func (g *graph) allVisited(path *pathAndPosition) bool {
	visited := map[[2]int]bool{}
	for coord := range *g {
		visited[coord] = false
	}
	for path.path != nil {
		visited[path.position] = true
		path = path.path
	}
	for _, v := range visited {
		if !v {
			return false
		}
	}
	return true
}
func (g *graph) connected(coords, coords2 [2]int) bool {
	x, y, x2, y2 := coords[0], coords[1], coords2[0], coords2[1]
	if (x != x2 && y != y2) || (x == x2 && y == y2) {
		return false
	}
	if x == x2 {
		lrg, sml := max(y, y2), min(y, y2)
		for sml != lrg {
			if lines[x][sml] == '.' {
				return false
			}
			sml++
		}
		return true
	}
	if y == y2 {
		lrg, sml := max(x, x2), min(x, x2)
		for sml != lrg {
			if lines[sml][y] == '.' {
				return false
			}
			sml++
		}
		return true
	}
	return false
}

func (g *graph) shouldVisit(path *pathAndPosition, x, y int) bool {
	if path.path == nil {
		return true
	}
	p := path.path
	for p != nil {
		if p.position[0] == x && p.position[1] == y {
			return false
		}
		p = p.path
	}
	return true
}

// returns if point is an intersection, and gives alignment parameter
func (f *field) isIntersection(x, y int) (bool, int) {
	if (*f)[x][y] == '.' {
		return false, 0
	}
	neighbors := [][2]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}
	val := 4
	valid := map[[2]int]bool{
		{x + 1, y}: true,
		{x - 1, y}: true,
		{x, y + 1}: true,
		{x, y - 1}: true,
	}
	for _, coords := range neighbors {
		if coords[0] < 0 || coords[0] >= len(*f) || coords[1] < 0 || coords[1] >= len((*f)[coords[0]]) || (*f)[coords[0]][coords[1]] == '.' {
			val--
			valid[coords] = false
		}
	}
	if val > 1 {
		if val == 2 {
			x1, y1, x2, y2 := -1, -1, -1, -1
			for _, coords := range neighbors {
				if valid[coords] {
					if x1 == -1 {
						x1, y1 = coords[0], coords[1]
					} else {
						x2, y2 = coords[0], coords[1]
						break
					}
				}
			}
			if x1 != x2 && y1 != y2 {
				return true, (x * y)
			} else {
				return false, -1
			}
		}
		return true, (x * y)
	}
	return false, 0
}

func part2() {
	g := graph{}
	content, _ := os.ReadFile("output.txt")
	lines = strings.Split(string(content), "\n")
	for i := range lines {
		if lines[i] == "" {
			continue
		}
		for j := range lines[i] {
			intersects, _ := lines.isIntersection(i, j)
			if intersects {
				n := node{x: i, y: j}
				g[[2]int{i, j}] = &n
			}
		}
		g[[2]int{38, 30}] = &node{x: 38, y: 30}
		g[[2]int{16, 36}] = &node{x: 16, y: 36}
	}
	for coords, n := range g {
		for coords2, n2 := range g {
			if n == n2 || !g.connected(coords, coords2) {
				continue
			}
			larger := 0
			if coords[0] == coords2[0] { // left & right
				diff := math.Abs(float64(coords[1] - coords2[1]))
				if coords[1] > coords2[1] {
					larger = 1
				}
				if larger == 1 {
					g[coords].left, g[coords].lEdge, g[coords2].right, g[coords2].rEdge = g[coords2], int(diff), g[coords], int(diff)
				} else {
					g[coords].right, g[coords].rEdge, g[coords2].left, g[coords2].lEdge = g[coords2], int(diff), g[coords], int(diff)
				}
			} else if coords[1] == coords2[1] { //up & down
				diff := math.Abs(float64(coords[0] - coords2[0]))
				if coords[0] > coords2[0] {
					larger = 1
				}
				if larger == 1 {
					g[coords].up, g[coords].uEdge, g[coords2].down, g[coords2].dEdge = g[coords2], int(diff), g[coords], int(diff)
				} else {
					g[coords].down, g[coords].dEdge, g[coords2].up, g[coords2].uEdge = g[coords2], int(diff), g[coords], int(diff)
				}

			}
		}
	}
	startPosition := pathAndPosition{nil, [2]int{16, 36}, '^', 0}
	queue := []pathAndPosition{startPosition}
	finished := []pathAndPosition{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(cur.toString())
		if g.allVisited(&cur) || cur.position == [2]int{38, 30} {
			finished = append(finished, cur)
			continue
		}
		node := g[cur.position]
		if node.lEdge > 0 && g.shouldVisit(&cur, node.left.x, node.left.y) {
			newPath := pathAndPosition{&cur, [2]int{node.left.x, node.left.y}, '<', (node.lEdge)}
			queue = append(queue, newPath)
		}
		if node.rEdge > 0 && g.shouldVisit(&cur, node.right.x, node.right.y) {
			newPath := pathAndPosition{&cur, [2]int{node.right.x, node.right.y}, '>', node.rEdge}
			queue = append(queue, newPath)

		}
		if node.uEdge > 0 && g.shouldVisit(&cur, node.up.x, node.up.y) {
			newPath := pathAndPosition{&cur, [2]int{node.up.x, node.up.y}, '^', node.uEdge}
			queue = append(queue, newPath)

		}
		if node.dEdge > 0 && g.shouldVisit(&cur, node.down.x, node.down.y) {
			newPath := pathAndPosition{&cur, [2]int{node.down.x, node.down.y}, 'v', node.dEdge}
			queue = append(queue, newPath)

		}
	}
	fmt.Println(g)

	fmt.Println(finished)
	for _, p := range finished {
		fmt.Println(p.toString())
	}

}

func part1() (int, error) {
	base := 0
	command, length := parse()
	command = append(command, make([]int, 100000)...)
	output := ""
	defer func() {
		sum := 0
		file, _ := os.Create("output.txt")
		file.WriteString(output)
		file.Close()

		content, _ := os.ReadFile("output.txt")
		var lines field = strings.Split(string(content), "\n")
		for i := range lines {
			if lines[i] == "" {
				continue
			}
			for j := range lines[i] {
				intersects, score := lines.isIntersection(i, j)
				if intersects {
					fmt.Println(i, j)
					sum += score
				}
			}
		}
		fmt.Println(sum)
	}()
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

			i += 2
		case 4: //output
			// if value1 == 10 {
			// 	output += "\n"
			// }
			output += string(value1)

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
	lines := [100]string{}
	line := ""
	for i := 0; i < 100; i++ {
		line += "."
	}
	for i := range lines {
		lines[i] = line
	}
	for i := range lines {
		fmt.Println(lines[i])
	}
	fmt.Println(lines)
	// fmt.Println(command, "f")
	if err != nil {
		// panic("problem, output is " + output)
		return -1, fmt.Errorf("bad output somehow")
	}
	return outputNum, nil

}
