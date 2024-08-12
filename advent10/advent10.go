package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type tile bool
type tiles [][]tile

func (t tiles) isVisible(ax, ay, bx, by int) bool {
	return false
}

func (t tiles) visible(acoords, bcoords [2]int) bool {
	if !t[bcoords[0]][bcoords[1]] {
		return false
	}
	xdist, ydist := bcoords[0]-acoords[0], bcoords[1]-acoords[1]
	xmagnitude, ymagnitude := xdist, ydist
	if xmagnitude < 0 {
		xmagnitude *= -1
	}
	if ymagnitude < 0 {
		ymagnitude *= -1
	}
	if xdist == 0 && ydist == 0 {
		return false
	}
	// fmt.Println(xdist, ydist)
	if xdist == 0 {
		if ydist > 0 {
			ydist = 1
		} else {
			ydist = -1
		}
		ymagnitude = 1
	}
	if ydist == 0 {
		if xdist > 0 {
			xdist = 1
		} else {
			xdist = -1
		}
		xmagnitude = 1
	}
	if ymagnitude == xmagnitude {
		ydist /= ymagnitude
		xdist /= xmagnitude
		xmagnitude /= xmagnitude
		ymagnitude /= ymagnitude
	}
	if ymagnitude != 0 && xmagnitude%ymagnitude == 0 {
		xmagnitude /= ymagnitude
		xdist /= ymagnitude
		ydist /= ymagnitude
		ymagnitude /= ymagnitude
	}
	if xmagnitude != 0 && ymagnitude%xmagnitude == 0 {
		ymagnitude /= xmagnitude
		xdist /= xmagnitude
		ydist /= xmagnitude
		xmagnitude /= xmagnitude
	}
	for i := 2; i < max(xmagnitude, ymagnitude); i++ {
		for xmagnitude%i == 0 && ymagnitude%i == 0 {
			xdist /= i
			ydist /= i
			xmagnitude /= i
			ymagnitude /= i
		}
	}
	// fmt.Println(xdist, ydist)

	x, y := acoords[0]+xdist, acoords[1]+ydist
	fmt.Println(x, y, xdist, ydist)
	for x != bcoords[0] || y != bcoords[1] {
		fmt.Println(x, y)
		if t[x][y] {
			fmt.Println(bcoords, "not visible from", acoords)
			return false
		}
		x += xdist
		y += ydist
	}
	fmt.Println(t[bcoords[0]][bcoords[1]], bcoords, "visible from", acoords, t[acoords[0]][acoords[1]])
	return true
}

func parse() tiles {
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
	lines := strings.Split(string(content), "\n")
	ret := make([][]tile, 0)
	for i, line := range lines {
		ret = append(ret, make([]tile, 0))
		for _, char := range strings.ReplaceAll(line, "\r", "") {
			if char == '.' {
				ret[i] = append(ret[i], false)
			} else {
				ret[i] = append(ret[i], true)
			}
		}
	}
	return ret
}

func verticalAngleFromOrigin(x, y int) float64 {
	// print(x, " ", y, " to ")
	x -= 16
	y -= 8
	// print(x, " ", y, "\n")
	if x == 0 && y > 0 {
		return math.Pi / 2.0
	} else if x == 0 && y < 0 {
		return 3.0 * math.Pi / 2.0
	} else if x < 0 && y == 0 {
		return 0.0
	} else if x > 0 && y == 0 {
		return math.Pi
	}
	var angle float64 = 0
	relativeQuadrant := 0

	if x > 0 && y > 0 {
		relativeQuadrant = 2
	} else if x < 0 && y < 0 {
		relativeQuadrant = 3
	} else if x < 0 && y > 0 {
		relativeQuadrant = 1
	}
	angle += (math.Pi * float64(relativeQuadrant) / 2.0)
	// fmt.Println(angle, relativeQuadrant, x, y)
	var floatx, floaty float64 = float64(x), float64(y)
	if floatx < 0 {
		floatx *= -1
	}
	if floaty < 0 {
		floaty *= -1
	}

	if relativeQuadrant%2 == 1 {
		// angle += math.Atan(float64(x) / float64(y))
		angle += math.Atan(floatx / floaty)
	} else {
		// angle += math.Atan(float64(y) / float64(y))
		angle += math.Atan(floaty / floatx)
	}
	// println(angle)
	return angle
}

func part2() {
	fmt.Println(verticalAngleFromOrigin(17, 8))
	fmt.Println(verticalAngleFromOrigin(15, 8))
	fmt.Println(verticalAngleFromOrigin(17, 0))
	tiles := parse()
	println(tiles[15][8])
	anglesByMagnitudes := map[float64][]float64{}
	anglesAndMagnitudesByCoordinates := make(map[[2]float64][2]int)

	for i, row := range tiles {
		for j, tile := range row {
			if tile && (i != 16 || j != 8) {
				magnitude := math.Sqrt(((float64(i) - 16) * (float64(i) - 16)) + ((float64(j) - 8) * (float64(j) - 8)))
				angle := verticalAngleFromOrigin(i, j)
				anglesAndMagnitudesByCoordinates[[2]float64{angle, magnitude}] = [2]int{i, j}
				if anglesByMagnitudes[angle] == nil {
					anglesByMagnitudes[angle] = []float64{magnitude}
				} else {
					anglesByMagnitudes[angle] = append(anglesByMagnitudes[angle], magnitude)
					slices.Sort(anglesByMagnitudes[angle])
				}
			}
		}
	}

	angles := make([]float64, 0)
	for key := range anglesByMagnitudes {
		angles = append(angles, key)
	}
	slices.Sort(angles)
	index := 0
	fmt.Println(angles)
	for i := 0; i < 200; {
		if len(anglesByMagnitudes[angles[index]]) > 0 {
			fmt.Println(anglesByMagnitudes[angles[index]], angles[index], anglesAndMagnitudesByCoordinates[[2]float64{angles[index], anglesByMagnitudes[angles[index]][0]}])
			// fmt.Println(anglesAndMagnitudesByCoordinates[[2]float64{angles[index], anglesByMagnitudes[angles[index]][0]}], " destroyed, number ", i)
			anglesByMagnitudes[angles[index]] = anglesByMagnitudes[angles[index]][1:]
			i++
		}
		index = (index + 1) % len(angles)

	}
	for len(anglesByMagnitudes[angles[index]]) == 0 {

		index = (index + 1) % len(angles)
	}

	fmt.Println(anglesByMagnitudes[angles[index]], angles[index], anglesAndMagnitudesByCoordinates[[2]float64{angles[index], anglesByMagnitudes[angles[index]][0]}])
}

func part1() {
	tiles := parse()
	maxim := 0
	maxCoords := [2]int{0, 0}
	for i, row := range tiles {
		for j, tile := range row {
			visibleCount := 0
			if !tile {
				continue
			}
			for k, row2 := range tiles {
				for l, tile2 := range row2 {
					// fmt.Printf("%d,%d,%d,%d\n", i, j, k, l)
					if !tile2 {
						continue
					}
					if tiles.visible([2]int{i, j}, [2]int{k, l}) {
						visibleCount++
					}
				}
			}
			fmt.Println(i, j, visibleCount)
			if visibleCount > maxim {
				maxim = visibleCount
				maxCoords[0], maxCoords[1] = i, j
			}
		}
	}
	fmt.Println(maxim, maxCoords, len(tiles[0]))
	fmt.Println(tiles)

}

func main() {
	part1()
	part2()
}
