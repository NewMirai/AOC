package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

type Octopus struct {
	energyLevel int
	hasFlashed  bool
}

func (o *Octopus) IncreaseEnergy() {
	o.energyLevel++
}

func (o *Octopus) Reset() {
	o.energyLevel = 0
}

type Position struct {
	x int
	y int
}

func (p *Position) GetAdjacentPositions() []Position {
	adj := make([]Position, 8)
	adj[0] = Position{x: p.x, y: p.y - 1}
	adj[1] = Position{x: p.x + 1, y: p.y - 1}
	adj[2] = Position{x: p.x + 1, y: p.y}
	adj[3] = Position{x: p.x + 1, y: p.y + 1}
	adj[4] = Position{x: p.x, y: p.y + 1}
	adj[5] = Position{x: p.x - 1, y: p.y + 1}
	adj[6] = Position{x: p.x - 1, y: p.y}
	adj[7] = Position{x: p.x - 1, y: p.y - 1}
	return adj
}

type Grid struct {
	grid         map[Position]*Octopus
	flashCounter int
}

func (g *Grid) Update() {
	for _, o := range g.grid {
		o.IncreaseEnergy()
	}
}

func (g *Grid) Display(nrow int, ncol int) {
	for i := 0; i < nrow; i++ {
		row := make([]int, ncol)
		for j := 0; j < ncol; j++ {
			row[j] = g.grid[Position{x: j, y: i}].energyLevel
		}
		fmt.Println(row)
	}
}

func (g *Grid) CheckSync(nrow int, ncol int) bool {
	n0 := 0
	for i := 0; i < nrow; i++ {
		for j := 0; j < ncol; j++ {
			el := g.grid[Position{x: j, y: i}].energyLevel
			if el == 0 {
				n0++
			}
		}
	}
	if n0 == 100 {
		return true
	}
	return false
}

func (g *Grid) ResetFlashed() {
	for _, o := range g.grid {
		if o.energyLevel > 9 {
			o.Reset()
		}
		o.hasFlashed = false
	}
}

func (o *Octopus) Flash(nFlashes *int, p *Position, g *Grid) {
	o.hasFlashed = true
	*nFlashes++
	adjPos := p.GetAdjacentPositions()
	for _, pos := range adjPos {
		if ao, ok := g.grid[pos]; ok {
			ao.IncreaseEnergy()
			if ao.energyLevel > 9 && !ao.hasFlashed {
				ao.Flash(nFlashes, &pos, g)
			}
		}
	}
}

func (g *Grid) Step() {
	g.Update()
	nFlashes := 0
	for p, o := range g.grid {
		if o.energyLevel > 9 && !o.hasFlashed {
			o.Flash(&nFlashes, &p, g)
		}
	}
	g.ResetFlashed()
	g.flashCounter += nFlashes
}

func Solve(s *string) int {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	n := len(lines)
	g := Grid{grid: make(map[Position]*Octopus, n*n), flashCounter: 0}
	for y, line := range lines {
		for x, level := range line {
			p := Position{x: x, y: y}
			energyLevel, _ := strconv.Atoi(string(level))
			o := Octopus{
				energyLevel: energyLevel,
				hasFlashed:  false,
			}
			g.grid[p] = &o
		}
	}
	step := 1
	// Part 1 for step < 100 {
	for step > 0 {
		g.Step()
		fmt.Println("After step", step)
		g.Display(n, n)
		if g.CheckSync(n, n) {
			break
		}
		step++
	}
	// Part 1 return g.flashCounter
	return step
}

// main function
func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println(Solve(&problem_input))
}
