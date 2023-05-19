package game

import (
	"fmt"
)

var round *Round

type Round struct {
	// Total words completed
	Total int
	// Total incorrect words
	Incorrect int
	// Score with per character points
	Score int
}

func GetRound() *Round {
	if round != nil {
		return round
	}

	fmt.Println("setting up new round")
	round = &Round{}
	return round
}

func (r *Round) Correct(w string) (int, int) {
	r.Total++
	r.Score += len(w)
	return r.Total, r.Score
}

func (r *Round) Wrong() int {
	r.Incorrect++
	return r.Incorrect
}

func (r *Round) FinalScore() int {
	return r.Score
}

func (r *Round) Reset() *Round {
	round = &Round{}
	return GetRound()
}
