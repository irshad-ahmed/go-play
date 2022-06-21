package rps

import (
	"math/rand"
	"time"
)

const (
	ROCK     = 0 // beats scissors. (scissors + 1) % 3 = 0
	PAPER    = 1 // beats rock. (rock + 1) % 3 = 1
	SCISSORS = 2 // beats paper. (paper + 1) % 3 = 2
)

type Round struct {
	Message        string `json:"message"`
	ComputerChoice string `json:"computerChoice"`
	RoundResult    string `json:"roundResult"`
}

var winMessage = []string{"Well Done", "Great Work", "You should buy a lottery ticket"}

var looseMessage = []string{"Better Luck Next Time", "Try again", "This is just not your day"}

var drawMessage = []string{"Nobody Wins", "Hey its a draw, Try again", "Great mind think alike"}

func PlayRound(playerValue int) Round {

	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := ""
	roundResult := ""
	var message string

	switch computerValue {
	case ROCK:
		computerChoice = "Computer chose ROCK"
	case PAPER:
		computerChoice = "Computer chose PAPER"
	case SCISSORS:
		computerChoice = "Computer chose SCISSORS"
	default:
	}

	messageInt := rand.Intn(3)
	if playerValue == computerValue {
		roundResult = "It's a draw"
		message = drawMessage[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		roundResult = "Player wins!"
		message = winMessage[messageInt]

	} else {
		roundResult = "Computer wins!"
		message = looseMessage[messageInt]
	}
	var result Round
	result.Message = message
	result.ComputerChoice = computerChoice
	result.RoundResult = roundResult
	return result

}
