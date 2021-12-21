package main

import (
	"fmt"
	"strings"
)

type Player struct {
	name  string
	pos   int
	score int
}

type deterministic_die int

func (d *deterministic_die) next() int {
	*d++
	if *d > 100 {
		*d %= 100
	}
	return int(*d)
}

func (p *Player) move(steps int) {
	new_pos := p.pos + steps
	for new_pos > 10 {
		new_pos %= 10
		if new_pos == 0 {
			new_pos = 10
		}
	}
	p.pos = new_pos
	p.score += p.pos
	fmt.Printf("score is now: %d\n", p.score)
}

func day21(puzzle_data []string) {
	/*
		for _, dataVal := range puzzle_data {
		}
	*/
	p1 := Player{
		name:  "one",
		score: 0,
	}
	p2 := Player{
		name:  "two",
		score: 0,
	}

	t := strings.Fields(puzzle_data[0])
	p1.pos = makeInt(t[len(t)-1])
	t = strings.Fields(puzzle_data[1])
	p2.pos = makeInt(t[len(t)-1])

	var d deterministic_die

	rolls := 0
	for p1.score < 1000 && p2.score < 1000 {
		for _, player := range []*Player{&p1, &p2} {
			move := 0
			fmt.Printf("Player %s rolled: ", player.name)
			for i := 0; i < 3; i++ {
				roll := d.next()
				move += roll
				rolls++
				fmt.Printf("%d ", roll)
			}
			player.move(move)
			if player.score >= 1000 {
				break
			}
		}
	}

	losing_score := p1.score
	if p2.score < losing_score {
		losing_score = p2.score
	}

	fmt.Printf("Day 21 (part 1): Losing score %d mul by rolls %d is %d\n", losing_score, rolls, losing_score*rolls)
}
