package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type reactionPart struct {
	label  string
	amount int
}

type reactions map[reactionPart][]reactionPart

func parse() reactions {
	reactions := reactions{}
	content, _ := os.ReadFile("input.txt")
	// content, _ := os.ReadFile("test.txt")
	// content, _ := os.ReadFile("test2.txt")
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		line = strings.ReplaceAll(line, "\r", "")
		if line == "" {
			break
		}
		io := strings.Split(line, " => ")
		fmt.Println(io)
		inputs := strings.Split(io[0], ", ")
		output := io[1]
		outputAmt, _ := strconv.Atoi(output[:strings.Index(output, " ")])
		outputLabel := output[strings.Index(output, " ")+1:]
		out := reactionPart{label: outputLabel, amount: outputAmt}
		is := []reactionPart{}
		for _, i := range inputs {
			inputAmt, _ := strconv.Atoi(i[:strings.Index(i, " ")])
			inputLabel := i[strings.Index(i, " ")+1:]
			is = append(is, reactionPart{label: inputLabel, amount: inputAmt})
		}

		reactions[out] = is
		fmt.Println(inputs, output)
	}
	return reactions
}

func main() {
	part1()
	part2driver()
}

type queueElement struct {
	label        string
	amountNeeded int
}

func part2driver() {
	i := 1122000
	for part2(i) < 1000000000000 {
		i++
	}
	fmt.Println(part2(1130000))
	fmt.Println(i)
}

func part2(x int) int {
	reactions := parse()
	amountsRemaining := map[string]int{}
	curOre := 0

	queue := make([]queueElement, 0)
	queue = append(queue, queueElement{label: "FUEL", amountNeeded: x})
	for len(queue) > 0 {
		cur := queue[0]
		fmt.Println(queue, curOre)
		queue = queue[1:]
		if cur.label == "ORE" {
			curOre += cur.amountNeeded
			continue
		}
		arcl := amountsRemaining[cur.label]
		println(arcl)
		if amountsRemaining[cur.label] > cur.amountNeeded {
			amountsRemaining[cur.label] -= cur.amountNeeded
			continue
		} else {
			cur.amountNeeded -= amountsRemaining[cur.label]
			amountsRemaining[cur.label] = 0
		}
		if possibleRe := reactions[reactionPart{label: cur.label, amount: cur.amountNeeded}]; possibleRe != nil {
			if len(possibleRe) > 1 {
				for _, re := range possibleRe {
					queue = append(queue, queueElement{label: re.label, amountNeeded: re.amount})
				}
			} else {
				if possibleRe[0].label == "ORE" {
					curOre += possibleRe[0].amount
				} else {
					queue = append(queue, queueElement{label: possibleRe[0].label, amountNeeded: possibleRe[0].amount})
				}
			}
		} else {
			for key, val := range reactions {
				if key.label == cur.label {
					// multiplyPart := 1
					// for x := key.amount; x < cur.amountNeeded; {
					// 	x += key.amount
					// 	multiplyPart += 1
					// }
					var multiplyPart int
					if cur.amountNeeded%key.amount == 0 {
						multiplyPart = cur.amountNeeded / key.amount
					} else {
						multiplyPart = (cur.amountNeeded / key.amount) + 1
					}
					fmt.Println(multiplyPart, (cur.amountNeeded/key.amount + 1))
					amountsRemaining[cur.label] += (key.amount * multiplyPart) - cur.amountNeeded
					for _, part := range val {
						if part.label == "ORE" {
							curOre += (part.amount * multiplyPart)
						} else {
							queue = append(queue, queueElement{label: part.label, amountNeeded: part.amount * multiplyPart})
						}
					}
					continue
				}
			}
		}
	}
	fmt.Println(amountsRemaining)
	return curOre

}

func part1() {
	reactions := parse()
	amountsRemaining := map[string]int{}
	curOre := 0

	queue := make([]queueElement, 0)
	queue = append(queue, queueElement{label: "FUEL", amountNeeded: 1})
	for len(queue) > 0 {
		cur := queue[0]
		fmt.Println(queue, curOre)
		queue = queue[1:]
		if cur.label == "ORE" {
			curOre += cur.amountNeeded
			continue
		}
		arcl := amountsRemaining[cur.label]
		println(arcl)
		if amountsRemaining[cur.label] > cur.amountNeeded {
			amountsRemaining[cur.label] -= cur.amountNeeded
			continue
		} else {
			cur.amountNeeded -= amountsRemaining[cur.label]
			amountsRemaining[cur.label] = 0
		}
		if possibleRe := reactions[reactionPart{label: cur.label, amount: cur.amountNeeded}]; possibleRe != nil {
			if len(possibleRe) > 1 {
				for _, re := range possibleRe {
					queue = append(queue, queueElement{label: re.label, amountNeeded: re.amount})
				}
			} else {
				if possibleRe[0].label == "ORE" {
					curOre += possibleRe[0].amount
				} else {
					queue = append(queue, queueElement{label: possibleRe[0].label, amountNeeded: possibleRe[0].amount})
				}
			}
		} else {
			for key, val := range reactions {
				if key.label == cur.label {
					// multiplyPart := 1
					// for x := key.amount; x < cur.amountNeeded; {
					// 	x += key.amount
					// 	multiplyPart += 1
					// }
					var multiplyPart int
					if cur.amountNeeded%key.amount == 0 {
						multiplyPart = cur.amountNeeded / key.amount
					} else {
						multiplyPart = (cur.amountNeeded / key.amount) + 1
					}
					fmt.Println(multiplyPart, (cur.amountNeeded/key.amount + 1))
					amountsRemaining[cur.label] += (key.amount * multiplyPart) - cur.amountNeeded
					for _, part := range val {
						if part.label == "ORE" {
							curOre += (part.amount * multiplyPart)
						} else {
							queue = append(queue, queueElement{label: part.label, amountNeeded: part.amount * multiplyPart})
						}
					}
					continue
				}
			}
		}
	}
	fmt.Println(amountsRemaining)
	fmt.Println(curOre)

}
