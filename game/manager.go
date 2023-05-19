package game

var round *Round

type Round struct {
	// Total words completed
	total int
	// Total incorrect words
	incorrect int
	// Score with per character points
	score int
}

func GetRound() *Round {
	if round != nil {
		return round
	}

	round = &Round{}
	return round
}

func (r *Round) Correct(w string) (int, int) {
	r.total++
	r.score += len(w)
	return r.total, r.score
}

func (r *Round) Incorrect() int {
	r.incorrect++
	return r.incorrect
}

func (r *Round) FinalScore() int {
	return r.score
}

func (r *Round) Reset() *Round {
	round = &Round{}
	return GetRound()
}
