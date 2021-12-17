package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(f string) string {
	data, _ := os.ReadFile(f)
	return string(data)
}

type Coordinate struct {
	x int
	y int
}

func (c *Coordinate) GetAdjacentCoordinates() []*Coordinate {
	adj := make([]*Coordinate, 4)
	adj[0] = &Coordinate{x: c.x, y: c.y - 1}
	adj[1] = &Coordinate{x: c.x + 1, y: c.y}
	adj[2] = &Coordinate{x: c.x, y: c.y + 1}
	adj[3] = &Coordinate{x: c.x - 1, y: c.y}
	return adj
}

type Vertex struct {
	cost  int
	pos   *Coordinate
	index int
}

type VertexQueue []*Vertex

func (vq VertexQueue) Len() int { return len(vq) }

func (vq VertexQueue) Less(i, j int) bool {
	// Return the lowest cost
	return vq[i].cost < vq[j].cost
}

func (vq VertexQueue) Swap(i, j int) {
	vq[i], vq[j] = vq[j], vq[i]
	vq[i].index = i
	vq[j].index = j
}

func (vq *VertexQueue) Push(x interface{}) {
	n := len(*vq)
	vertex := x.(*Vertex)
	vertex.index = n
	*vq = append(*vq, vertex)
}

func (vq *VertexQueue) Pop() interface{} {
	old := *vq
	n := len(*vq)
	vertex := old[n-1]
	old[n-1] = nil    // avoid memory leak
	vertex.index = -1 // for safety
	*vq = old[0 : n-1]
	return vertex
}

func (vq *VertexQueue) update(vertex *Vertex, pos *Coordinate, cost int) {
	vertex.pos = &Coordinate{x: pos.x, y: pos.y}
	vertex.cost = cost
	heap.Fix(vq, vertex.index)
}

func Min(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func MinPathSum(grid [][]int, H int, W int) int {
	dp := make([][]int, H)
	for i := 0; i < H; i++ {
		dp[i] = make([]int, W)
		if i == 0 {
			dp[i][0] = grid[0][0]
		} else {
			dp[i][0] = dp[i-1][0] + grid[i][0]
		}
	}
	for j := 1; j < W; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < H; i++ {
		for j := 1; j < W; j++ {
			dp[i][j] = grid[i][j] + Min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[H-1][W-1] - dp[0][0]
}

func MinPathSum2(grid [][]int, H int, W int) int {
	total := 0
	vq := make(VertexQueue, 0)
	start := &Coordinate{x: 0, y: 0}
	target := &Coordinate{x: W - 1, y: H - 1}
	heap.Init(&vq)
	initV := &Vertex{cost: 0, pos: start}
	heap.Push(&vq, initV)
	visited := map[Coordinate]struct{}{
		*start: struct{}{},
	}
	for {
		v := heap.Pop(&vq).(*Vertex)
		cost, pos := v.cost, v.pos
		if pos.x == target.x && pos.y == target.y {
			total = cost
			break
		}
		for _, apos := range pos.GetAdjacentCoordinates() {
			_, ok := visited[*apos]
			if 0 <= apos.y && apos.y < H && 0 <= apos.x && apos.x < W && !ok {
				newV := &Vertex{
					cost: grid[apos.y][apos.x] + cost,
					pos:  apos,
				}
				heap.Push(&vq, newV)
				visited[*apos] = struct{}{}
			}
		}
	}
	return total
}

func Repeat(grid [][]int) [][]int {
	H := len(grid)
	W := len(grid[0])
	newGrid := make([][]int, H)
	for i := 0; i < H; i++ {
		newGrid[i] = make([]int, W)
		for j := 0; j < W; j++ {
			newValue := grid[i][j] + 1
			if newValue > 9 {
				newValue = 1
			}
			newGrid[i][j] = newValue
		}
	}
	return newGrid
}

func Solve(s *string) (total int) {
	lines := strings.Split(strings.TrimSpace(*s), "\n")
	H := len(lines)
	W := len(lines[0])
	grid := make([][]int, H)
	for i, line := range lines {
		grid[i] = make([]int, W)
		for j := 0; j < W; j++ {
			cost, _ := strconv.Atoi(string(line[j]))
			grid[i][j] = cost
		}
	}
	// part 1
	fmt.Println("Part 1 using dp: ", MinPathSum(grid, H, W))

	r2 := Repeat(grid)
	r3 := Repeat(r2)
	r4 := Repeat(r3)
	r5 := Repeat(r4)
	bigGrid := make([][]int, H)
	for i := 0; i < H; i++ {
		bigGrid[i] = make([]int, 5*W)
		for j := 0; j < W; j++ {
			for k := 0; k < 5; k++ {
				switch k {
				case 0:
					bigGrid[i][j+H*k] = grid[i][j]
				case 1:
					bigGrid[i][j+H*k] = r2[i][j]
				case 2:
					bigGrid[i][j+H*k] = r3[i][j]
				case 3:
					bigGrid[i][j+H*k] = r4[i][j]
				case 4:
					bigGrid[i][j+H*k] = r5[i][j]
				}
			}
		}
	}
	r2c := Repeat(bigGrid)
	r3c := Repeat(r2c)
	r4c := Repeat(r3c)
	r5c := Repeat(r4c)
	bigGridComplete := make([][]int, 5*len(bigGrid))
	for i := 0; i < 5*H; i++ {
		bigGridComplete[i] = make([]int, len(bigGrid[0]))
	}
	for k := 0; k < 5; k++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				bigGridComplete[i][j+H*k] = bigGrid[i][j+H*k]
			}
		}
	}
	for k := 0; k < 5; k++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				bigGridComplete[i+H][j+H*k] = r2c[i][j+H*k]
			}
		}
	}
	for k := 0; k < 5; k++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				bigGridComplete[i+2*H][j+H*k] = r3c[i][j+H*k]
			}
		}
	}
	for k := 0; k < 5; k++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				bigGridComplete[i+3*H][j+H*k] = r4c[i][j+H*k]
			}
		}
	}
	for k := 0; k < 5; k++ {
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				bigGridComplete[i+4*H][j+H*k] = r5c[i][j+H*k]
			}
		}
	}
	// part 2
	H = len(bigGridComplete)
	W = len(bigGridComplete[0])
	total = MinPathSum2(bigGridComplete, H, W)
	return total
}

func main() {
	problem_input := ReadInput("input.txt")
	fmt.Println("Part 2 using heap: ", Solve(&problem_input))
}
