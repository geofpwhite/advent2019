package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type moon struct {
	px, py, pz, vx, vy, vz int
}
type axisValues struct {
	m1p, m1v, m2p, m2v, m3p, m3v, m4p, m4v int
}

func (m moons) toAxisValues(axis rune) axisValues {
	switch axis {
	case 'x':
		return axisValues{
			m[0].px, m[0].vx,
			m[1].px, m[1].vx,
			m[2].px, m[2].vx,
			m[3].px, m[3].vx,
		}
	case 'y':
		return axisValues{
			m[0].py, m[0].vy,
			m[1].py, m[1].vy,
			m[2].py, m[2].vy,
			m[3].py, m[3].vy,
		}
	case 'z':
		return axisValues{
			m[0].pz, m[0].vz,
			m[1].pz, m[1].vz,
			m[2].pz, m[2].vz,
			m[3].pz, m[3].vz,
		}
	default:
		panic("bad axis")
	}

}

type moons []moon

func main() {
	part2()
}

func part1() {
	moons := parse()
	fmt.Println(moons)
	for i := 0; i < 1000; i++ {
		moons.applyGravity()
		moons.applyVelocity()
		fmt.Println(moons)
	}
	energy := 0
	for _, m := range moons {
		p, v := 0, 0
		p += int(math.Abs(float64(m.px)))
		p += int(math.Abs(float64(m.py)))
		p += int(math.Abs(float64(m.pz)))
		v += int(math.Abs(float64(m.vx)))
		v += int(math.Abs(float64(m.vy)))
		v += int(math.Abs(float64(m.vz)))
		energy += p * v
	}
	fmt.Println(energy)
}
func part2() {
	// _moons := parse()
	// fmt.Println(moons)

	fmt.Println(lcm(84, 252, 33))
	fmt.Println(gcd(5, 10))
	//if it goes back to a previous state, it must cycle; maybe it cycles for each moon and we need to take the lcm of all of them
	axes := "xyz"
	axis := 0
	indices := []int{}
	for axis < 3 {
		// prevMoons := map[moon]int{}
		// prevAxisValues := map[axisValues]int{}
		moons := parse()
		// start := parse()
		prevAxisValues := map[axisValues]int{moons.toAxisValues(rune(axes[axis])): 1}
		i := 0
		for { //i := 0; i < 1000; i++ {
			moons.applyGravity()
			moons.applyVelocity()

			if prevAxisValues[moons.toAxisValues(rune(axes[axis]))] > 0 {
				indices = append(indices, i+1)
				axis++
				break
			}

			// prevMoons[moons[curMoonIndex]]++
			prevAxisValues[moons.toAxisValues(rune(axes[axis]))]++
			i++
			// fmt.Println(_moons)
		}

	}
	fmt.Println(indices)
	fmt.Println(lcm(indices...))

}

func lcm(nums ...int) int { // lcm(x,y) = xy/gcd(x,y)
	curNum := nums[0]
	nextNumIndex := 1
	for nextNumIndex < len(nums) {
		curNum = curNum * nums[nextNumIndex] / gcd(curNum, nums[nextNumIndex])
		nextNumIndex++
	}
	return curNum
}
func gcd(x, y int) int {
	gcd := 1
	curDivisor := 2
	if y%x == 0 {
		return x
	}
	if x%y == 0 {
		return y
	}
	for curDivisor < max(x, y) && x > 1 && y > 1 {
		for x%curDivisor == 0 && y%curDivisor == 0 {
			x /= curDivisor
			y /= curDivisor
			gcd *= curDivisor
		}
		curDivisor++
	}
	return gcd
}

func parse() moons {
	moons := moons{}
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if !strings.Contains(line, "<") {
			continue
		}
		positionStrings := strings.Split(line, ",")
		positionStrings[0] = strings.TrimSpace(positionStrings[0][1:])
		positionStrings[2] = strings.TrimSpace(positionStrings[2][:len(positionStrings[2])-1])
		positionStrings[1] = strings.TrimSpace(positionStrings[1])
		fmt.Println(positionStrings)
		for i := range positionStrings {
			positionStrings[i] = strings.ReplaceAll(positionStrings[i], ">", "")
		}

		px, _ := strconv.Atoi(positionStrings[0][2:])
		py, _ := strconv.Atoi(positionStrings[1][2:])
		pz, _ := strconv.Atoi(positionStrings[2][2:])
		fmt.Println(px, py, pz)
		m := moon{px: px, py: py, pz: pz}
		moons = append(moons, m)
	}
	return moons
}

func (m moons) applyGravity() {
	for i := range m {
		for j := i + 1; j < len(m); j++ {
			m1, m2 := m[i], m[j]
			if m1.px > m2.px {
				m[i].vx--
				m[j].vx++
			} else if m1.px < m2.px {
				m[j].vx--
				m[i].vx++
			}
			if m1.py > m2.py {
				m[i].vy--
				m[j].vy++
			} else if m1.py < m2.py {
				m[j].vy--
				m[i].vy++
			}
			if m1.pz > m2.pz {
				m[i].vz--
				m[j].vz++
			} else if m1.pz < m2.pz {
				m[j].vz--
				m[i].vz++
			}
		}
	}
}

func (m moons) applyVelocity() {
	for i := range m {
		m[i].px += m[i].vx
		m[i].py += m[i].vy
		m[i].pz += m[i].vz
	}
}
