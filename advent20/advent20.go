package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type node struct {
	x, y      int
	neighbors map[*node]int
}
type coords [2]int

type graph map[coords]*node

func connected(x1, y1, x2, y2 int, lines []string, g graph) bool {
	if x1 != x2 && y1 != y2 {
		return false
	}
	if x1 == x2 {
		for i := min(y1, y2) + 1; i < max(y1, y2); i++ {
			if lines[x1][i] != '.' || g[coords{x1, i}] != nil {
				return false
			}
		}
		return true
	} else {
		for i := min(x1, x2) + 1; i < max(x1, x2); i++ {
			if lines[i][y1] != '.' || g[coords{i, y1}] != nil {
				return false
			}
		}
		return true
	}

}
func parse() (graph, map[string][]coords) {
	g := make(graph)
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	defer func() {
		f, _ := os.Create("output.txt")
		ary := make([]string, 118)
		for i, line := range lines {
			for j := range line {
				if g[coords{i, j}] == nil {
					ary[i] += " "
				} else {
					ary[i] += "."
				}
			}
		}
		f.WriteString(strings.Join(ary, "\n"))
	}()

	for i, line := range lines {
		for j, char := range line {
			if char == '.' && cornerOrIntersection(lines, i, j) {
				n := &node{i, j, make(map[*node]int)}
				g[coords{i, j}] = n
			}
		}
	}
	portals := map[string][]coords{}
	length := len(lines)
	/* 	for i, char := range lines[0] {
	   		if char != ' ' {
	   			str := string(lines[0][i]) + string(lines[1][i])
	   			portals[str] = []coords{{2, i}}
	   		}
	   	}
	   	for i, char := range lines[0] {
	   		if char != ' ' {
	   			str := string(lines[0][i]) + string(lines[1][i])
	   			if portals[str] == nil {
	   				portals[str] = []coords{{2, i}}

	   			} else {
	   				portals[str] = append(portals[str], coords{2, i})
	   			}
	   		}
	   	}
	   	for i, char := range lines[length-2] {
	   		if char != ' ' {
	   			str := string(lines[length-3][i]) + string(lines[length-2][i])
	   			if portals[str] == nil {
	   				portals[str] = []coords{{length - 3, i}}
	   			} else {
	   				portals[str] = append(portals[str], coords{length - 3, i})
	   			}
	   		}
	   	} */
	uCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, line := range lines {
		for j, char := range line {
			if strings.Contains(uCase, string(char)) && j < length {
				if strings.Contains(uCase, string(line[j+1])) {
					//horizontal
					str := line[j : j+2]
					if j > 0 && line[j-1] == '.' {
						if portals[str] == nil {
							portals[str] = []coords{{i, j - 1}}
						} else {
							portals[str] = append(portals[str], coords{i, j - 1})
						}

					} else if j < length-3 && line[j+2] == '.' {
						if portals[str] == nil {
							portals[str] = []coords{{i, j + 2}}
						} else {
							portals[str] = append(portals[str], coords{i, j + 2})
						}

					}
				} else if i < length-2 && strings.Contains(uCase, string(lines[i+1][j])) {
					//vertical
					str := string(lines[i][j]) + string(lines[i+1][j])
					if i > 0 && lines[i-1][j] == '.' {
						if portals[str] == nil {
							portals[str] = []coords{{i - 1, j}}
						} else {
							portals[str] = append(portals[str], coords{i - 1, j})
						}

					} else if i < length-3 && lines[i+2][j] == '.' {
						if portals[str] == nil {
							portals[str] = []coords{{i + 2, j}}
						} else {
							portals[str] = append(portals[str], coords{i + 2, j})
						}

					}

				}
			}
		}
	}
	for key, val := range portals {
		fmt.Println(key, val, g[val[0]])
		if len(val) == 2 {
			g[val[0]].neighbors[g[val[1]]] = 1
			g[val[1]].neighbors[g[val[0]]] = 1
		}

	}
	for _, n := range g {
		for _, n2 := range g {
			if n2 == n || n.neighbors[n2] > 0 {
				continue
			}
			if connected(n.x, n.y, n2.x, n2.y, lines, g) {
				// fmt.Println(n.x, n.y, n2.x, n2.y)
				n.neighbors[n2] = int(max(math.Abs(float64(n.x-n2.x)), math.Abs(float64(n.y-n2.y))))
				n2.neighbors[n] = int(max(math.Abs(float64(n.x-n2.x)), math.Abs(float64(n.y-n2.y))))
			}
		}
	}
	return g, portals
}
func cornerOrIntersection(lines []string, x, y int) bool {
	if lines[x][y] != '.' {
		return false
	}
	neighbors := [4][2]int{
		{x + 1, y},
		{x - 1, y},
		{x, y + 1},
		{x, y - 1},
	}
	valid := [4]bool{}
	num := 0
	for i, coords := range neighbors {
		if lines[coords[0]][coords[1]] != '#' {
			valid[i] = true
			num++
		}

		if lines[coords[0]][coords[1]] != '#' && lines[coords[0]][coords[1]] != '.' {
			return true
		}
	}
	if num != 2 && num != 0 {
		return true
	}
	i1, i2 := -1, -1
	for i := range neighbors {
		if valid[i] {
			if i1 == -1 {
				i1 = i
			} else {
				i2 = i
				break
			}
		}
	}
	if (min(i1, i2) == 0 && max(i1, i2) == 1) || (min(i1, i2) == 2 && max(i1, i2) == 3) {
		return false
	}
	return true
}

func main() {
	part1()
	part2()
}
func part1() {
	g, portals := parse()
	// for _, n := range g {
	// 	fmt.Println(n)
	// }
	start := g[portals["AA"][0]]
	end := g[portals["ZZ"][0]]
	distances := map[*node]int{}
	for _, n := range g {
		distances[n] = -1
	}
	type qE struct {
		node  *node
		steps int
	}

	queue := []qE{{start, 0}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if cur.steps < distances[cur.node] || distances[cur.node] == -1 {
			distances[cur.node] = cur.steps
			for n, edge := range cur.node.neighbors {
				queue = append(queue, qE{n, cur.steps + edge})
			}
		}
	}
	// for k, v := range g {
	// 	fmt.Print(k, "->")
	// 	for k2, v := range v.neighbors {
	// 		fmt.Print(k2.x, ",", k2.y, " ", v, "     ")
	// 	}
	// 	fmt.Println()
	// }
	// for k, v := range distances {
	// 	fmt.Println(k, v)
	// }
	fmt.Println(len(distances), len(g))
	fmt.Println(distances[end])
}

func part2() {
	g, p := parse()

	type qE struct {
		node         *node
		steps, level int
	}
	distances := make(map[*node]map[int]int)
	for _, n := range g {
		distances[n] = make(map[int]int)
	}
	start := qE{g[p["AA"][0]], 0, 0}
	queue := []qE{start}
	reverseAccessPortals := map[[2]coords]int{}
	for lbl, coordAry := range p {
		if lbl == "AA" || lbl == "ZZ" {
			continue
		}
		c1, c2 := coordAry[0], coordAry[1]
		if ((c1[0] == 30 || c1[0] == 84) && (c1[1] > 31 && c1[1] < 86)) ||
			((c1[1] == 30 || c1[1] == 86) && (c1[0] > 31 && c1[0] < 86)) {
			reverseAccessPortals[[2]coords{c1, c2}] = 1 //inner to outer
			reverseAccessPortals[[2]coords{c2, c1}] = -1
			fmt.Println("inner 2 outer", c1, c2)
		} else {
			reverseAccessPortals[[2]coords{c1, c2}] = -1
			reverseAccessPortals[[2]coords{c2, c1}] = 1
			fmt.Println(" outer 2 inner", c1, c2)

		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		fmt.Println(len(queue), cur.level)
		if cur.node.x == g[p["ZZ"][0]].x && cur.node.y == g[p["ZZ"][0]].y && cur.level == 0 {
			println(cur.steps)
			return
		}
		if cur.steps < distances[cur.node][cur.level] || distances[cur.node][cur.level] == 0 {
			distances[cur.node][cur.level] = cur.steps
			for n, edge := range cur.node.neighbors {
				newLevel := cur.level + reverseAccessPortals[[2]coords{{cur.node.x, cur.node.y}, {n.x, n.y}}]
				if newLevel < 0 || newLevel > 100 {
					continue
				}
				queue = append(queue, qE{n, cur.steps + edge, newLevel})
			}
		}
	}
	fmt.Println(distances[g[p["ZZ"][0]]][0])
}
