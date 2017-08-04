package main

import (
	"testing"
)

func TestMakePlayer(t *testing.T) {
	testPlayer := makePlayer("test")
	testPlay := <-testPlayer

	found := false
	for _, m := range moves {
		if testPlay.move == m {
			found = true
		}
	}
	if !found {
		t.Errorf("Unexpected move, %s", testPlay.move)
	}

	if testPlay.name != "test" {
		t.Errorf("Unexpected player name, %s", testPlay.name)
	}
}

func TestMakeJudge(t *testing.T) {
	winner := make(chan play)
	loser := make(chan play)

	go func() {
		winner <- play{name: "winner", move: rock}
		loser <- play{name: "loser", move: scissors}
	}()

	judge := makeJudge(winner, loser)
	testResult := <-judge

	if testResult.winner != "winner" {
		t.Errorf("winner player should have won but got %s", testResult.winner)
	}
}

func TestWinnerIs(t *testing.T) {
	rockPlay := play{name: "rock", move: rock}
	paperPlay := play{name: "paper", move: paper}
	scissorsPlay := play{name: "scissors", move: scissors}

	allMatches := []result{
		{rockPlay, rockPlay, "no one"},
		{rockPlay, paperPlay, "paper"},
		{rockPlay, scissorsPlay, "rock"},
		{paperPlay, rockPlay, "paper"},
		{paperPlay, paperPlay, "no one"},
		{paperPlay, scissorsPlay, "scissors"},
		{scissorsPlay, rockPlay, "rock"},
		{scissorsPlay, paperPlay, "scissors"},
		{scissorsPlay, scissorsPlay, "no one"},
	}

	for _, match := range allMatches {
		winner := winnerIs(match.play1, match.play2)
		if winner != match.winner {
			t.Errorf("%s should have beat %s but got %s", match.play1.move, match.play2.move, winner)
		}
	}
}
