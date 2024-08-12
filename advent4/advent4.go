package main

import "strconv"

func main() {
	//240298-784956
	sum := 0
	validFlag := true
	for i := 240298; i < 784956; i++ {
		validFlag = true
		adjacentFlag := false
		adjacentCount := map[int]int{
			0: 0,
			1: 0,
			2: 0,
			3: 0,
			4: 0,
			5: 0,
			6: 0,
			7: 0,
			8: 0,
			9: 0,
		}
		str := strconv.Itoa(i)
		curDigit, lastDigit := 0, 0
		for _, char := range str {
			curDigit, _ = strconv.Atoi(string(char))
			if curDigit < lastDigit {
				validFlag = false
				break
			}
			if curDigit == lastDigit {
				adjacentCount[curDigit] += 2
			}

			lastDigit = curDigit
		}
		if validFlag {
			for _, val := range adjacentCount {
				if val == 2 {
					adjacentFlag = true
				}
			}
			if adjacentFlag {
				sum++
			}
		}
	}
	println(sum)
}
