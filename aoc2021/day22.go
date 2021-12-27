package main

import (
	"fmt"
	"regexp"
)

type Cube struct {
	minX int
	maxX int
	minY int
	maxY int
	minZ int
	maxZ int
}

func day22(puzzle_data []string) {
	// brute force, only works for part 1
	cubemap := map[string]bool{}
	inputRegex := regexp.MustCompile(`(on|off) x\=(-?\d+)..(-?\d+),y\=(-?\d+)..(-?\d+),z\=(-?\d+)..(-?\d+)`)
	for _, dataVal := range puzzle_data {
		minX := 0
		minY := 0
		minZ := 0
		maxX := 0
		maxY := 0
		maxZ := 0
		cmd := false
		matches := inputRegex.FindStringSubmatch(dataVal)
		if len(matches) > 0 {
			if matches[1] == "on" {
				cmd = true
			}
			minX = makeInt(matches[2])
			maxX = makeInt(matches[3])
			minY = makeInt(matches[4])
			maxY = makeInt(matches[5])
			minZ = makeInt(matches[6])
			maxZ = makeInt(matches[7])
		}
		if minZ >= -50 && minY >= -50 && minX >= -50 && maxZ <= 50 && maxY <= 50 && maxX <= 50 {
			for z := minZ; z <= maxZ; z++ {
				for y := minY; y <= maxY; y++ {
					for x := minX; x <= maxX; x++ {
						key := fmt.Sprintf("%d,%d,%d", x, y, z)
						if cmd {
							cubemap[key] = cmd
						} else {
							delete(cubemap, key)
						}
					}
				}
			}
		}
	}
	fmt.Printf("Day 22 (part 1):The number of cubes turned on is %d\n", len(cubemap))

	cubes := []Cube{}
	for _, dataVal := range puzzle_data {
		cmd := false
		matches := inputRegex.FindStringSubmatch(dataVal)
		if len(matches) > 0 {
			if matches[1] == "on" {
				cmd = true
			}
			cube := Cube{
				minX: makeInt(matches[2]),
				maxX: makeInt(matches[3]),
				minY: makeInt(matches[4]),
				maxY: makeInt(matches[5]),
				minZ: makeInt(matches[6]),
				maxZ: makeInt(matches[7]),
			}
			cubes = merge(cubes, cube, cmd)
		}
	}
	fmt.Printf("Day 22 (part 2):The number of cubes turned on is %d\n", len(cubes))
}

func (c Cube) intersects(c2 Cube) bool {
	if c.minX >= c2.minX && c.minX <= c2.maxX {
	}
	if c.minY >= c2.minY && c.minY <= c2.maxY {
	}
	if c.minZ >= c2.minZ && c.minZ <= c2.maxZ {
	}
	return false
}

func merge(cubes []Cube, cube Cube, onoff bool) []Cube {
	newCubes := []Cube{}

	return newCubes
}
