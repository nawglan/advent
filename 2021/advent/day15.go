package main

import (
	"container/heap"
	"fmt"
	"math"
)

var (
	grid []int
)

type Item struct {
	value    int // pos
	priority int // distance
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

// we want lowest distance first
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, pos int, distance int) {
	item.value = pos
	item.priority = distance
	heap.Fix(pq, item.index)
}

func dijkstra(start int) ([]int, []int) {
	seen := make([]bool, gridWidth*gridLength)
	dist := make([]int, gridWidth*gridLength)
	prev := make([]int, gridWidth*gridLength)
	pq := PriorityQueue{}
	heap.Init(&pq)

	// init distance
	for i := range dist {
		dist[i] = math.MaxInt
	}
	dist[start] = 0

	// add start to the priority queue
	item := &Item{
		value:    start,
		priority: 0,
	}
	heap.Push(&pq, item)

	for pq.Len() > 0 {
		curr_item := heap.Pop(&pq).(*Item)
		seen[curr_item.value] = true
		for _, neighbor := range neighbors(curr_item.value) {
			if seen[neighbor] {
				continue
			}
			newDist := dist[curr_item.value] + grid[neighbor]
			if newDist < dist[neighbor] {
				prev[neighbor] = curr_item.value
				dist[neighbor] = newDist
				new_item := &Item{
					value:    neighbor,
					priority: newDist,
				}
				heap.Push(&pq, new_item)
			}
		}
	}

	return dist, prev
}

func findShortestPath(start, end int) []int {
	dist, prev := dijkstra(0)

	rev_path := []int{}
	if dist[end] == math.MaxInt {
		return rev_path
	}
	for at := end; at != start; at = prev[at] {
		rev_path = append(rev_path, at)
	}
	rev_path = append(rev_path, start)

	path := []int{}
	for i := len(rev_path) - 1; i >= 0; i-- {
		path = append(path, rev_path[i])
	}

	return path
}

func dumpPath(shortest_path []int) {
	step := 0
	for i := 0; i < len(grid); i++ {
		x := i % gridWidth
		if step <= len(shortest_path)-1 && i == shortest_path[step] {
			fmt.Printf(". ")
			step++
		} else {
			fmt.Printf("%d ", grid[i])
		}
		if x == gridWidth-1 {
			fmt.Println()
		}
	}
}

func day15(puzzle_data []string) {
	grid = makeIntGrid(puzzle_data, len(puzzle_data[0]), len(puzzle_data))

	shortest_path := findShortestPath(0, len(grid)-1)

	lowest_risk := 0
	for _, pos := range shortest_path {
		lowest_risk += grid[pos]
	}
	lowest_risk -= grid[0] // we get the fist step for free
	fmt.Printf("Day 15 (part 1): Lowest total risk is %d\n", lowest_risk)

	riskLevelTemplate := map[int][]int{}
	for riskLevel_y := 0; riskLevel_y < 5; riskLevel_y++ {
		for riskLevel_x := 0; riskLevel_x < 5; riskLevel_x++ {
			idx := riskLevel_y + riskLevel_x
			if _, found := riskLevelTemplate[idx]; !found {
				riskLevelTemplate[idx] = make([]int, len(grid))
				for i := 0; i < len(grid); i++ {
					riskLevelTemplate[idx][i] = grid[i] + idx
					if riskLevelTemplate[idx][i] > 9 {
						riskLevelTemplate[idx][i] %= 10
						riskLevelTemplate[idx][i]++
					}
				}
			}
		}
	}

	newGridWidth := gridWidth * 5
	newGridLength := gridLength * 5
	full_grid := make([]int, newGridWidth*newGridLength)
	for y_risk := 0; y_risk < 5; y_risk++ {
		for x_risk := 0; x_risk < 5; x_risk++ {
			for i := 0; i < len(riskLevelTemplate[y_risk+x_risk]); i++ {
				full_x := (x_risk * gridWidth) + (i % gridWidth)
				full_y := (y_risk * gridLength) + (i / gridWidth)
				idx := full_x + full_y*newGridWidth
				full_grid[idx] = riskLevelTemplate[y_risk+x_risk][i]
			}
		}
	}

	grid = full_grid
	gridWidth = newGridWidth
	gridLength = newGridLength

	shortest_path = findShortestPath(0, len(grid)-1)
	//dumpPath(shortest_path)
	lowest_risk = 0
	for _, pos := range shortest_path {
		lowest_risk += grid[pos]
	}
	lowest_risk -= grid[0] // we get the fist step for free
	fmt.Printf("Day 15 (part 2): Lowest total risk is %d\n", lowest_risk)
}
