package game

import (
	"github.com/alien_invasion/game/engine"
	_map "github.com/alien_invasion/game/map"
)

const (
	mapFilename string = "map_test"
)

type Game struct {
	*engine.Engine
}

func New(numAliens uint64) *Game {
	return &Game{
		engine.New(numAliens, _map.New(mapFilename)),
	}
}

func (g *Game) Run() {
	g.Run()
}
