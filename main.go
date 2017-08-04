package main

import (
	"fmt"
	"math/rand"
	"time"
)

type move string

type play struct {
	move move
	name string
}

type result struct {
	play1, play2 play
	winner       string
}

const (
	rock     move = "rock"
	paper    move = "paper"
	scissors move = "scissors"
)

var moves = [3]move{rock, paper, scissors}

func main() {
	// Wouldn't want to get the same "random" numbers every time
	rand.Seed(time.Now().Unix())

	player1 := makePlayer("Amon")
	player2 := makePlayer("Belinda")
	judge := makeJudge(player1, player2)

	totals := make(map[string]int)
	for i := 0; i < 10000; i++ {
		r := <-judge
		fmt.Printf("%s throws %s\n", r.play1.name, r.play1.move)
		fmt.Printf("%s throws %s\n", r.play2.name, r.play2.move)
		fmt.Printf("%s wins!\n", r.winner)
		totals[r.winner]++
	}
	fmt.Println(totals)
}

func makePlayer(name string) <-chan play {
	out := make(chan play)
	go func() {
		for { // ever
			i := rand.Intn(len(moves))
			out <- play{name: name, move: moves[i]}
		}
	}()
	return out
}

func makeJudge(player1, player2 <-chan play) <-chan result {
	out := make(chan result)
	go func() {
		for { // ever
			p1 := <-player1
			p2 := <-player2
			out <- result{play1: p1, play2: p2, winner: winnerIs(p1, p2)}
		}
	}()
	return out
}

// wins[move] should tell you what move it BEATS
var wins = map[move]move{
	rock:     scissors,
	paper:    rock,
	scissors: paper,
}

func winnerIs(play1, play2 play) string {
	if play1.move == play2.move {
		return "no one"
	}
	if wins[play1.move] == play2.move {
		return play1.name
	}
	return play2.name
}
