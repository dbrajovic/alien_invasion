package engine

import "github.com/alien_invasion/game/types"

type Engine struct {
	aliens  []types.Alien
	gameMap Map
}

func New(aliens []types.Alien, gameMap Map) *Engine {
	return &Engine{aliens, gameMap}
}

func (e *Engine) Run() {
	for {
		//	check if game should stop

		//	move aliens

		//	fight aliens

		//	remove destroyed cities (if any)
	}
}
