package game

import (
	"log"
	"os"

	"github.com/alien_invasion/game/engine"
	_map "github.com/alien_invasion/game/map"
)

const (
	mapFilename string = "game_map"
)

type Game struct {
	*engine.Engine
}

func New(numAliens uint64) *Game {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal("cannot open file:", "err=", err)
	}

	return &Game{
		engine.New(
			numAliens,
			_map.New(cwd+"/"+mapFilename),
		),
	}
}
