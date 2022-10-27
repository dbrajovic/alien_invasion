package engine

import "github.com/alien_invasion/game/types"

const (
	maxTravels uint64 = 10000
)

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
		if e.isDone() {
			return
		}

		//	move aliens

		//	fight aliens

		//	remove destroyed cities (if any)
	}
}

func (e *Engine) isDone() bool {
	return e.areAliensDead() || e.areAliensMaxedOutOnTravels()
}

func (e *Engine) areAliensDead() bool {
	return len(e.aliens) == 0
}

func (e *Engine) areAliensMaxedOutOnTravels() bool {
	for _, alien := range e.aliens {
		if alien.Steps < maxTravels {
			return false
		}
	}

	return true
}
