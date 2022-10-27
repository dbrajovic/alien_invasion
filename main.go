package main

import (
	"fmt"
	"log"

	"github.com/alien_invasion/game"
)

func main() {
	var input uint64

	if _, err := fmt.Scanf("%d", &input); err != nil {
		log.Fatal("failed to read input: %w", err)
	}

	//	init game
	game.New().Run()

	//	run game
}
