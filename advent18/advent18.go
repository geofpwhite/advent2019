package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"strings"
)

type node struct {
	value                 rune
	x, y                  int
	up, down, left, right *node
}

type queueElement struct {
	node          *node
	steps         int
	keysCollected map[rune]bool
}

type coords [2]int
type graph map[coords]*node

func cornerOrIntersection(lines []string, x, y int) bool {
	if lines[x][y] == '#' {
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
func connected(x1, y1, x2, y2 int, lines []string) bool {
	if x1 != x2 && y1 != y2 {
		return false
	}
	if x1 == x2 {
		for i := min(y1, y2) + 1; i < max(y1, y2); i++ {
			if lines[x1][i] != '.' {
				return false
			}
		}
		return true
	} else {
		for i := min(x1, x2) + 1; i < max(x1, x2); i++ {
			if lines[i][y1] != '.' {
				return false
			}
		}
		return true
	}

}

func parse() (g graph) {
	g = make(graph)
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		for j, char := range line {
			if char != '#' {
				if char != '.' {
					n := &node{x: i, y: j, value: char}
					g[coords{i, j}] = n
				} else {
					if cornerOrIntersection(lines, i, j) {
						n := &node{x: i, y: j, value: '.'}
						g[coords{i, j}] = n
					}
				}
			}
		}
	}
	for co, n := range g {
		for co2, n2 := range g {
			if n == n2 {
				continue
			}
			// 0,0
			if connected(co[0], co[1], co2[0], co2[1], lines) {
				distance := int(max(math.Abs(float64(co[0]-co2[0])), math.Abs(float64(co[1]-co2[1]))))

				if co[0] == co2[0] {
					if co[1] < co2[1] && (n.up == nil || int(math.Abs(float64(co[1]-n.up.y))) > distance) {
						n.up = n2
						n2.down = n
					} else if co[1] > co2[1] && (n.down == nil || int(math.Abs(float64(co[1]-n.down.y))) > distance) {
						n.down = n2
						n2.up = n
					}
				} else {
					if co[0] < co2[0] && (n.right == nil || int(math.Abs(float64(co[0]-n.right.x))) > distance) {
						n.right = n2
						n2.left = n
					} else if co[0] > co2[0] && (n.left == nil || int(math.Abs(float64(co[0]-n.left.x))) > distance) {
						n.left = n2
						n2.right = n
					}
				}
			}
		}
	}
	return g
}
func parse2() (g graph) {
	g = make(graph)
	content, _ := os.ReadFile("input2.txt")
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		for j, char := range line {
			if char != '#' {
				if char != '.' {
					n := &node{x: i, y: j, value: char}
					g[coords{i, j}] = n
				} else {
					if cornerOrIntersection(lines, i, j) {
						n := &node{x: i, y: j, value: '.'}
						g[coords{i, j}] = n
					}
				}
			}
		}
	}
	for co, n := range g {
		for co2, n2 := range g {
			if n == n2 {
				continue
			}
			// 0,0
			if connected(co[0], co[1], co2[0], co2[1], lines) {
				distance := int(max(math.Abs(float64(co[0]-co2[0])), math.Abs(float64(co[1]-co2[1]))))

				if co[0] == co2[0] {
					if co[1] < co2[1] && (n.up == nil || int(math.Abs(float64(co[1]-n.up.y))) > distance) {
						n.up = n2
						n2.down = n
					} else if co[1] > co2[1] && (n.down == nil || int(math.Abs(float64(co[1]-n.down.y))) > distance) {
						n.down = n2
						n2.up = n
					}
				} else {
					if co[0] < co2[0] && (n.right == nil || int(math.Abs(float64(co[0]-n.right.x))) > distance) {
						n.right = n2
						n2.left = n
					} else if co[0] > co2[0] && (n.left == nil || int(math.Abs(float64(co[0]-n.left.x))) > distance) {
						n.left = n2
						n2.right = n
					}
				}
			}
		}
	}
	return g
}

func sortAlphabetically(str string) string {
	ary := []rune{}
	for _, char := range str {
		ary = append(ary, char)
	}
	for i := range ary {
		for j := i; j > 0 && ary[j] < ary[j-1]; j-- {
			ary[j], ary[j-1] = ary[j-1], ary[j]
		}
	}
	newString := ""
	for _, char := range ary {
		newString += string(char)
	}
	return newString
}

const lCase string = "abcdefghijklmnopqrstuvwxyz"

func part1() {
	g := parse()

	// for coord, node := range g {
	// 	fmt.Println(node, coord, string(node.value))
	// }
	fmt.Println(len(g))
	var startNode *node
	for _, n := range g {
		if n.value == '@' {
			startNode = n
			break
		}
	}
	shortestPathsToNodeByUnlockedKeys := map[*node]map[string]int{}

	startQueue := queueElement{node: startNode, steps: 0, keysCollected: make(map[rune]bool)}
	queue := []queueElement{startQueue}
	lowerNodes := ""
	for _, n := range g {
		if strings.Contains(lCase, string(n.value)) {
			lowerNodes += string(n.value)
		}
	}
	for len(queue) > 0 {
		// elem := queue[len(queue)-1]
		// queue = queue[:len(queue)-1]
		elem := queue[0]
		queue = queue[1:]
		n := elem.node
		unlocked := ""
		if strings.Contains(lCase, string(n.value)) {
			elem.keysCollected[n.value] = true
		}
		for key := range elem.keysCollected {
			unlocked += string(key)
		}
		unlocked = sortAlphabetically(unlocked)
		// if unlocked != "" {
		// 	fmt.Println(unlocked)
		// }
		if shortestPathsToNodeByUnlockedKeys[n] == nil {
			shortestPathsToNodeByUnlockedKeys[n] = make(map[string]int)
		}
		if sh := shortestPathsToNodeByUnlockedKeys[n][unlocked]; (sh == 0) || sh > elem.steps {
			shortestPathsToNodeByUnlockedKeys[n][unlocked] = elem.steps
		} else {
			continue
		}

		if n.value != '.' && n.value != '@' {
			if strings.Contains(lCase, string(n.value)) {
				elem.keysCollected[n.value] = true
				neighbors := []*node{n.up, n.down, n.left, n.right}
				for _, ne := range neighbors {
					if ne != nil {
						distance := max(math.Abs(float64(ne.x-n.x)), math.Abs(float64(ne.y-n.y)))
						newKeys := make(map[rune]bool)
						maps.Copy(newKeys, elem.keysCollected)
						newElement := queueElement{node: ne, steps: elem.steps + int(distance), keysCollected: newKeys}
						queue = append(queue, newElement)
					}
				}
			} else {
				lowerCase := strings.ToLower(string(n.value))
				if !elem.keysCollected[rune(lowerCase[0])] {
					continue
				} else {
					neighbors := []*node{n.up, n.down, n.left, n.right}
					for _, ne := range neighbors {
						if ne != nil {
							distance := max(math.Abs(float64(ne.x-n.x)), math.Abs(float64(ne.y-n.y)))
							newKeys := make(map[rune]bool)
							maps.Copy(newKeys, elem.keysCollected)
							newElement := queueElement{node: ne, steps: elem.steps + int(distance), keysCollected: newKeys}
							queue = append(queue, newElement)
						}
					}
				}
			}
		} else {
			neighbors := []*node{n.down, n.left, n.right, n.up}

			for _, ne := range neighbors {
				if ne != nil {
					distance := max(math.Abs(float64(ne.x-n.x)), math.Abs(float64(ne.y-n.y)))
					newKeys := make(map[rune]bool)
					maps.Copy(newKeys, elem.keysCollected)
					newElement := queueElement{node: ne, steps: elem.steps + int(distance), keysCollected: newKeys}
					queue = append(queue, newElement)
				}
			}

		}
	}

	ma := -1
	for _, value := range shortestPathsToNodeByUnlockedKeys {
		if value[lCase] < ma || ma == -1 {
			ma = value[lCase]
		}
	}
	fmt.Println(ma)
}

type robotQueuePositions struct {
	topLeft, topRight, bottomLeft, bottomRight *node
	unlocked                                   map[rune]bool
	steps                                      int
}

func (g *graph) bestMove(rqp robotQueuePositions) robotQueuePositions {
	newUnlocked := make(map[rune]bool)
	maps.Copy(newUnlocked, rqp.unlocked)
	newRqp := robotQueuePositions{rqp.topLeft, rqp.topRight, rqp.bottomLeft, rqp.bottomRight, newUnlocked, rqp.steps}
	nodes := []*node{rqp.topLeft, rqp.topRight, rqp.bottomLeft, rqp.bottomRight}
	distances := [4]map[rune]int{
		make(map[rune]int),
		make(map[rune]int),
		make(map[rune]int),
		make(map[rune]int),
	}
	for _, char := range lCase {
		if !rqp.unlocked[char] {
			for i, n := range nodes {
				_, distances[i][char] = g.shortestPathToKeyIfPossible(n, char, rqp.unlocked)
			}
		}
	}
	minDist := -1
	minChar := 'a'
	quadrant := -1
	for i, ma := range distances {
		for char, dist := range ma {
			if minDist == -1 || (dist != -1 && dist < minDist) {
				minDist, minChar = dist, char
				quadrant = i
			}
		}
	}
	for _, n := range *g {
		if n.value == minChar {
			switch quadrant {
			case 0:
				newRqp.topLeft = n
			case 1:
				newRqp.topRight = n
			case 2:
				newRqp.bottomLeft = n
			case 3:
				newRqp.bottomRight = n
			}
			newRqp.unlocked[minChar] = true
			break
		}
	}
	newRqp.steps = minDist + rqp.steps
	return newRqp
}

const uCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (g *graph) shortestPathToKeyIfPossible(n *node, key rune, unlocked map[rune]bool) (bool, int) {
	if unlocked[key] {
		return false, -1
	}
	start := queueElement{node: n, steps: 0, keysCollected: unlocked}
	queue := []queueElement{start}
	dists := map[*node]int{}
	keyNode := start.node
	for len(queue) > 0 {
		cur := queue[0]
		n := cur.node
		queue = queue[1:]
		if strings.Contains(uCase, string(n.value)) && !unlocked[rune(strings.ToLower(string(n.value))[0])] {
			continue
		}
		if n.value == key {
			keyNode = n
		}
		neighbors := []*node{cur.node.left, cur.node.right, cur.node.up, cur.node.down}
		for _, ne := range neighbors {
			if ne == nil {
				continue
			}
			distance := max(math.Abs(float64(ne.x-n.x)), math.Abs(float64(ne.y-n.y)))
			newUnlocked := make(map[rune]bool)
			maps.Copy(newUnlocked, cur.keysCollected)
			if ne != nil && (dists[ne] <= 0 || dists[ne] > int(distance)+cur.steps) {
				queue = append(queue, queueElement{ne, int(distance) + cur.steps, newUnlocked})
				dists[ne] = int(distance) + cur.steps
				fmt.Println(dists[ne])
			}
		}
	}

	fmt.Println(dists[keyNode])
	return dists[keyNode] > 0, dists[keyNode]
}

func part2() {
	g := parse2()

	// for coord, node := range g {
	// 	fmt.Println(node, coord, string(node.value))
	// }
	fmt.Println(len(g))
	var startNodes []*node = make([]*node, 0)
	for _, n := range g {
		if n.value == '@' {
			startNodes = append(startNodes, n)
		}
	}
	// shortestPathsToRobotPositionsByUnlockedKeys := map[[4]*node]map[string]int{}

	// startQueue := queueElement{node: startNode, steps: 0, keysCollected: make(map[rune]bool)}
	startQueue := robotQueuePositions{startNodes[0], startNodes[1], startNodes[2], startNodes[3], make(map[rune]bool), 0}
	lowerNodes := ""
	for _, n := range g {
		if strings.Contains(lCase, string(n.value)) {
			lowerNodes += string(n.value)
		}
	}
	// fmt.Println(len(queue))
	// elem := queue[len(queue)-1]
	// queue = queue[:len(queue)-1]
	elem := startQueue
	for len(elem.unlocked) < 26 {
		// fmt.Println(elem.unlocked)
		elem = g.bestMove(elem)
		fmt.Println(elem.steps)
	}
	fmt.Println(elem)
	fmt.Println(string(99))

}
func main() {
	part1()
	part2()
}
