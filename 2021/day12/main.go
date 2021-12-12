package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

type AjacencyList struct {
	values map[string][]string
}

type Cave struct {
	name    string
	visited map[string]struct{}
	twice   string
}

func Solve(s *string) (npaths int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	a := AjacencyList{make(map[string][]string)}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		left, right := parts[0], parts[1]
		if _, ok := a.values[left]; ok {
			a.values[left] = append(a.values[left], right)
		} else {
			a.values[left] = []string{right}
		}
		if _, ok := a.values[right]; ok {
			a.values[right] = append(a.values[right], left)
		} else {
			a.values[right] = []string{left}
		}

	}
	initialCave := Cave{name: "start", visited: make(map[string]struct{})}
	initialCave.visited["start"] = struct{}{}
	Caves := make([]Cave, 0)
	Caves = append(Caves, initialCave)
	for len(Caves) > 0 {
		cave := Caves[0]
		Caves = Caves[1:]
		if cave.name == "end" {
			npaths++
			continue
		}
		for _, el := range a.values[cave.name] {
			if _, ok := cave.visited[el]; !ok {
				copyVisited := make(map[string]struct{})
				for k, v := range cave.visited {
					copyVisited[k] = v
				}
				if strings.ToLower(el) == el {
					copyVisited[el] = struct{}{}
				}
				Caves = append(Caves, Cave{
					name:    el,
					visited: copyVisited,
					twice:   "",
				})
			} else {
				if cave.twice == "" && cave.name != "start" && cave.name != "end" {
					Caves = append(Caves, Cave{
						name:    el,
						visited: cave.visited,
						twice:   el,
					})
				}
			}
		}

	}
	return
}

func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
