package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type coord struct {
	x, y          int
	next          *coord
	steps         int
	nextDirection string
}

type path struct {
	curCoord *coord
}

func parse() []path {
	// content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test1.txt")
	content, _ := os.ReadFile("test2.txt")
	lines := strings.Split(string(content), "\n")
	path1s, path2s := strings.Split(lines[0], ","), strings.Split(lines[1], ",")
	path1, path2 := &path{&coord{x: 0, y: 0, next: nil, steps: 0}}, &path{&coord{x: 0, y: 0, next: nil, steps: 0}}
	p1c, p2c := path1.curCoord, path2.curCoord
	for _, dir := range path1s {
		nextCoord := coord{x: p1c.x, y: p1c.y, steps: p1c.steps}
		p1c.nextDirection = dir
		println(dir)
		switch dir[0] {
		case 'U':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.y += num
			nextCoord.steps += num
		case 'D':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.y -= num
			nextCoord.steps += num
		case 'R':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.x += num
			nextCoord.steps += num
		case 'L':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.x -= num
			nextCoord.steps += num
		default:
			println(dir)
			panic("ah")
		}
		p1c.next = &nextCoord
		p1c = p1c.next
	}
	for _, dir := range path2s {
		nextCoord := coord{x: p2c.x, y: p2c.y, steps: p2c.steps}
		println(dir)
		switch dir[0] {
		case 'U':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.y += num
			nextCoord.steps += num
		case 'D':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.y -= num
			nextCoord.steps += num
		case 'R':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.x += num
			nextCoord.steps += num
		case 'L':
			num, _ := strconv.Atoi(dir[1:])
			nextCoord.x -= num
			nextCoord.steps += num
		default:

			panic("ah")
		}
		p2c.next = &nextCoord
		p2c = p2c.next
	}
	p2c.next = &coord{0, 0, nil, 0, ""}
	p1c.next = &coord{0, 0, nil, 0, ""}
	return []path{*path1, *path2}
}

func main() {
	paths := parse()
	coord11, coord12 := *paths[0].curCoord, *paths[0].curCoord.next
	coord21, coord22 := *paths[1].curCoord, *paths[1].curCoord.next
	distances := []int{}
	fmt.Println(coord11, coord12, coord21, coord22)
	steps := []int{}

	for coord12.next != nil {
		coord22 = *paths[1].curCoord.next
		for coord22.next != nil {
			if intersecting, distance, _steps := intersect(coord11, coord12, coord21, coord22); intersecting {
				distances = append(distances, distance)
				steps = append(steps, _steps)
			}
			coord21 = coord22
			coord22 = *coord22.next

		}
		coord11 = coord12
		coord12 = *coord12.next
	}
	fmt.Println(slices.Min(distances))
	fmt.Println(slices.Min(steps))
	fmt.Println(steps)
}

// returns whether or not they intersect, the manhattan distance, and the total steps
func intersect(coord11, coord12, coord21, coord22 coord) (bool, int, int) {
	//return false,-1 if lines are parallel
	if (coord11.x == coord12.x && coord21.x == coord22.x) ||
		(coord11.y == coord12.y && coord21.y == coord22.y) ||
		(coord11.x != coord12.x && coord11.y != coord12.y) ||
		(coord21.x != coord22.x && coord21.y != coord22.y) {
		return false, -1, -1
	}
	vertical, horizontal := [2]coord{}, [2]coord{}
	if coord11.x == coord12.x {
		vertical[0], vertical[1] = coord11, coord12
		horizontal[0], horizontal[1] = coord21, coord22
	} else {
		vertical[0], vertical[1] = coord21, coord22
		horizontal[0], horizontal[1] = coord11, coord12
	}
	if vertical[0].y > vertical[1].y {
		vertical[0], vertical[1] = vertical[1], vertical[0]
	}
	if horizontal[0].x > horizontal[1].x {
		horizontal[0], horizontal[1] = horizontal[1], horizontal[0]
	}
	if (horizontal[0].x < vertical[0].x && horizontal[1].x > vertical[0].x) &&
		(vertical[0].y < horizontal[0].y && vertical[1].y > horizontal[0].y) {
		intersectionPoint := [2]int{vertical[0].x, horizontal[0].y}
		if intersectionPoint[0] < 0 {
			intersectionPoint[0] *= -1
		}
		if intersectionPoint[1] < 0 {
			intersectionPoint[1] *= -1
		}
		fmt.Println(coord11, coord12, "intersects with", coord21, coord22, "at", intersectionPoint)

		if vertical[1].steps < vertical[0].steps {
			vertical[0], vertical[1] = vertical[1], vertical[0]
		}

		if horizontal[1].steps < horizontal[0].steps {
			horizontal[0], horizontal[1] = horizontal[1], horizontal[0]
		}
		distanceC11 := vertical[0].y - intersectionPoint[1]
		if distanceC11 < 0 {
			distanceC11 *= -1
		}
		// distanceC12 := vertical[1].y - intersectionPoint[1]
		// if distanceC12 < 0 {
		// 	distanceC12 *= -1
		// }
		distanceC21 := horizontal[0].x - intersectionPoint[0]
		if distanceC21 < 0 {
			distanceC21 *= -1
		}
		// distanceC22 := horizontal[1].x - intersectionPoint[0]
		// if distanceC22 < 0 {
		// 	distanceC22 *= -1
		// }
		steps := horizontal[0].steps + vertical[0].steps + distanceC11 + distanceC21

		return true, intersectionPoint[0] + intersectionPoint[1], steps
	}
	return false, -1, -1
}
