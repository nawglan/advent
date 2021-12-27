package main

import (
	"fmt"
	"strings"
)

var (
	p1_start int
	p2_start int
	seen     = map[string][2]int64{}
)

type Player struct {
	pos   int
	score int
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
}

type deterministic_die int

func (d *deterministic_die) next() int {
	sum := 0
	for i := 0; i < 3; i++ {
		*d++
		if *d == 101 {
			*d = 1
		}
		sum += int(*d)
	}
	return sum
}

type Game struct {
	players    [2]Player
	rolls_left int
	turn       int // which player's turn it is
}

func start_game() (g Game) {
	g.players[0] = Player{
		pos:   p1_start,
		score: 0,
	}
	g.players[1] = Player{
		pos:   p2_start,
		score: 0,
	}
	g.turn = 0
	g.rolls_left = 3

	return g
}

func day21(puzzle_data []string) {
	// parse input data
	t := strings.Fields(puzzle_data[0])
	p1_start = makeInt(t[len(t)-1])
	t = strings.Fields(puzzle_data[1])
	p2_start = makeInt(t[len(t)-1])

	game := start_game()

	var shared_die deterministic_die
	player1 := Player{
		pos:   p1_start,
		score: 0,
	}
	player2 := Player{
		pos:   p2_start,
		score: 0,
	}

	rolls := 0
	for player1.score < 1000 && player2.score < 1000 {
		for _, player := range []*Player{&player1, &player2} {
			move := shared_die.next()
			rolls += 3
			player.move(move)
			player.score += player.pos
			if player.score >= 1000 {
				break
			}
		}
	}

	losing_score := player1.score
	if player2.score < losing_score {
		losing_score = player2.score
	}

	fmt.Printf("Day 21 (part 1): Losing score %d mul by rolls %d is %d\n", losing_score, rolls, losing_score*rolls)

	game = start_game()

	p1_wins, p2_wins := play(game)
	num_universes := p1_wins
	if p2_wins > p1_wins {
		num_universes = p2_wins
	}
	fmt.Printf("Day 21 (part 2): Num universes for the winner in more universes is %d\n", num_universes)
}

func play(g Game) (p1_win, p2_win int64) {
	key := fmt.Sprint(g)
	if res, ok := seen[key]; ok {
		return res[0], res[1]
	}

	if g.rolls_left == 0 {
		g.players[g.turn].score += g.players[g.turn].pos
		if g.players[g.turn].score > 20 {
			if g.turn == 0 {
				return 1, 0
			} else {
				return 0, 1
			}
		}
		if g.turn == 0 {
			g.turn = 1
		} else {
			g.turn = 0
		}
		g.rolls_left = 3
	}
	for roll := 1; roll <= 3; roll++ {
		g_copy := Game{
			players:    [2]Player{g.players[0], g.players[1]},
			rolls_left: g.rolls_left - 1,
			turn:       g.turn,
		}
		g_copy.players[g.turn].move(roll)
		one, two := play(g_copy)
		p1_win += one
		p2_win += two
	}

	seen[key] = [2]int64{p1_win, p2_win}

	return p1_win, p2_win
}
