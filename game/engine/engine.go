package engine

import (
	"github.com/alien_invasion/game/types"
)

const (
	maxTravels uint64 = 10000
)

type Engine struct {
	aliens  []*types.Alien
	gameMap Map
}

func New(aliens uint64, gameMap Map) *Engine {
	return &Engine{
		spawnRandomAliens(aliens, gameMap),
		gameMap,
	}
}

func (e *Engine) Run() {
	defer e.gameMap.Display()

	for {
		//	check if game should stop
		if e.isDone() {
			return
		}

		//	move aliens
		e.moveAliens()

		//	fight aliens
		destroyed := e.aliensFight()

		//	remove destroyed cities (if any)
		e.removeCities(destroyed...)
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
		if alien.Travels < maxTravels {
			return false
		}
	}

	return true
}

func (e *Engine) removeCities(cities ...types.City) {
	for _, city := range cities {
		e.gameMap.RemoveCity(city)
	}
}

func (e *Engine) moveAliens() {
	for _, alien := range e.aliens {
		alien.Location = e.gameMap.RandomNeighbourCity(alien.Location)
		alien.Travels++
	}
}

func (e *Engine) move(alien types.Alien) {

}

func (e *Engine) aliensFight() []types.City {

	aliensByCities := make(map[types.City][]*types.Alien)
	for _, alien := range e.aliens {
		existing := aliensByCities[alien.Location]
		existing = append(existing, alien)
		aliensByCities[alien.Location] = existing
	}

	var (
		destroyedCities []types.City
		deadAliens      []*types.Alien
	)

	for city, aliens := range aliensByCities {
		if len(aliens) > 1 {
			destroyedCities = append(destroyedCities, city)
			deadAliens = append(deadAliens, aliens...)

			//	print
		}
	}

	for _, a := range deadAliens {
		for i, aa := range e.aliens {
			if a.Name == aa.Name {
				e.aliens = append(e.aliens[:i], e.aliens[i+1:]...)
			}
		}
	}

	return destroyedCities
}
