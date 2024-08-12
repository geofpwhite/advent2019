package main

import (
	"fmt"
	"os"
	"strings"
)

type object struct {
	label  string
	orbits map[string]*object
}
type objectPart2 struct {
	label     string
	neighbors map[string]*objectPart2
}
type queueObject struct {
	*objectPart2
	steps int
}

func parse2() (objects map[string]*objectPart2) {
	objects = make(map[string]*objectPart2)
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		labels := strings.Split(line, ")")
		if len(labels) == 1 {
			break
		}
		if objects[labels[0]] == nil {
			newObject := &objectPart2{label: labels[0], neighbors: make(map[string]*objectPart2)}
			objects[labels[0]] = newObject
		}
		if objects[labels[1]] == nil {
			newObject := &objectPart2{label: labels[1], neighbors: make(map[string]*objectPart2)}
			objects[labels[1]] = newObject
		}
		objects[labels[1]].neighbors[labels[0]] = objects[labels[0]]
		objects[labels[0]].neighbors[labels[1]] = objects[labels[1]]
	}
	return
}

func parse() (objects map[string]*object) {
	objects = make(map[string]*object)
	content, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		labels := strings.Split(line, ")")
		if len(labels) == 1 {
			break
		}
		if objects[labels[0]] == nil {
			newObject := &object{label: labels[0], orbits: make(map[string]*object)}
			objects[labels[0]] = newObject
		}
		if objects[labels[1]] == nil {
			newObject := &object{label: labels[1], orbits: make(map[string]*object)}
			objects[labels[1]] = newObject
		}
		objects[labels[1]].orbits[labels[0]] = objects[labels[0]]
	}
	return
}
func main() {
	{
		objects := parse()
		sum := 0
		for _, obj := range objects {
			queue := []*object{obj}
			for len(queue) > 0 {
				poppedObject := queue[0]
				queue = queue[1:]
				for _, orbit := range poppedObject.orbits {
					sum++
					queue = append(queue, orbit)
				}
			}
		}
	}
	objects1 := parse()
	objects := parse2()
	start := objects["YOU"]
	for label := range objects1["YOU"].orbits {
		start = objects[label]
		break
	}
	dest := objects["SAN"]
	for label := range objects1["SAN"].orbits {
		dest = objects[label]
		break
	}
	visited := make(map[string]bool)
	lengths := make(map[string]int)
	queue := []*queueObject{&queueObject{start, 0}}
	var cur *queueObject
	for len(queue) > 0 {
		cur = queue[0]
		visited[cur.label] = true
		lengths[cur.label] = cur.steps
		queue = queue[1:]
		if cur.objectPart2 == dest {
			fmt.Println(cur.steps)
			break
		} else {
			for _, neighbor := range cur.neighbors {
				if !visited[neighbor.label] || lengths[neighbor.label] > cur.steps+1 {
					newQueueObject := queueObject{neighbor, cur.steps + 1}
					queue = append(queue, &newQueueObject)
				}
			}
		}
	}

}
