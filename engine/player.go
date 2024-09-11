package engine

import (
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms/googleai"
)

type Player struct {
	id         string
	brain      *googleai.GoogleAI
	color      string
	moveNumber int
	timeLeft   int
}

func (p Player) init(color string, apiKey string) *Player {
	brain, err := googleai.New(Ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}

	time := os.Getenv("GAMETIME_EACH_HALF")
	gameTime, _ := strconv.Atoi(time)

	id := uuid.New()

	return &Player{
		id:         id.String()[:7],
		brain:      brain,
		color:      color,
		moveNumber: 0,
		timeLeft:   gameTime,
	}
}
