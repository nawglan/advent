package main

import (
	"fmt"
	"math"
	"strings"
	p "advent/util/parse"
)

func day17(puzzle_data []string) {
	puzzle_data[0] = strings.Replace(puzzle_data[0], "target area: x=", "", 1)
	target_area := strings.Split(puzzle_data[0], ",")
	target_area[1] = strings.Replace(target_area[1], " y=", "", 1)

	target_x := strings.Split(target_area[0], "..")
	target_y := strings.Split(target_area[1], "..")

	target_minx := p.MakeInt(target_x[0])
	target_maxx := p.MakeInt(target_x[1])
	target_miny := p.MakeInt(target_y[0])
	target_maxy := p.MakeInt(target_y[1])

	max_y_hit := target_miny

	x, y := 0, 0 // start position

	seen := map[string]bool{}

	orig_yvel := target_miny * target_maxx
	for orig_yvel <= int(math.Abs(float64(target_miny))) {
		orig_xvel := 1
		for orig_xvel <= target_maxx*2 {
			yvel := orig_yvel
			xvel := orig_xvel
			x = 0
			y = 0
			max_y := 0
			step := 1
			for y >= target_miny && x <= target_maxx {
				// calc new x and new y
				x += xvel
				y += yvel
				if max_y < y {
					max_y = y
				}

				if xvel != 0 {
					if xvel > 0 {
						xvel -= 1
					} else {
						xvel += 1
					}
				}
				yvel--

				if y >= target_miny && y <= target_maxy && x >= target_minx && x <= target_maxx {
					seen[fmt.Sprintf("%d,%d", orig_xvel, orig_yvel)] = true
					if max_y > max_y_hit {
						max_y_hit = max_y
					}
					break
				}
				step++
			}
			orig_xvel++
		}
		orig_yvel++
	}

	fmt.Printf("Day 17 (part 1): The max y achieved is %d\n", max_y_hit)
	fmt.Printf("Day 17 (part 2): The count of initial velocities needed to hit the target is %d\n", len(seen))
}
