package main

import (
	"fmt"
	"log"

	"github.com/alien_invasion/game"
)

func main() {
	game.New(inputNumAliens()).Run()
}

func inputNumAliens() uint64 {
	var input uint64

	for {
		if _, err := fmt.Scanf("%d", &input); err != nil {
			log.Fatal("failed to read input: %w", err)
		}

		if input == 0 {
			fmt.Println("number must be greater than 0")

			continue
		}

		return input
	}
}
