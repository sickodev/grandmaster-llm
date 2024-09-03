package engine

import (
	"log"
	"os"
	"strconv"
	"time"
)

func StartDefault() {
	val, err := strconv.Atoi(os.Getenv("TOTAL_GAME_TIME"))

	if err != nil {
		log.Fatal("Failed setting up total game time")
	}

	totalTime := time.Duration(val) * time.Second

	for totalTime != 0 {

	}

}
