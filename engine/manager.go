package engine

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
)

var Ctx = context.Background()

func StartDefault() {

	val, err := strconv.Atoi(os.Getenv("TOTAL_GAME_TIME"))

	if err != nil {
		log.Fatal("Failed setting up total game time")
	}

	totalTime := time.Duration(val) * time.Second

	apiKey := os.Getenv("GEMINI_API_KEY")

	//Load Players
	white := Player{}
	white = *white.init("white", apiKey)
	println(white.id)

	black := Player{}
	black = *black.init("black", apiKey)
	println(black.id)

	log.Println("Game Time: " + totalTime.String())
	SetGame()
}

func SetGame() {
	gameId := uuid.New().String()[:16]

	rootDir, _ := os.Getwd()
	println(rootDir)
	fileName := fmt.Sprintf("%s.txt", gameId)
	filepath := rootDir + "/results/" + fileName
	print(filepath)
	_, err := os.Create(filepath)
	if err != nil {
		log.Panic("Error creating results file")
	}
}
