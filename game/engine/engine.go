package engine

import (
	"log"

	"github.com/alien_invasion/game/types"
)

const (
	maxTravels uint64 = 10000
)

// Engine is the main module that runs the game to completion
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
		if e.isDone() {
			return
		}

		e.moveAliens()

		destroyed := e.aliensFight()

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

func (e *Engine) aliensFight() []types.City {
	deadAliens, destroyedCities := getCasualties(e.groupAliensByCity())

	e.removeAliens(deadAliens)

	return destroyedCities
}

func (e *Engine) groupAliensByCity() map[types.City][]*types.Alien {
	aliensByCities := make(map[types.City][]*types.Alien)
	for _, alien := range e.aliens {
		existing := aliensByCities[alien.Location]
		existing = append(existing, alien)
		aliensByCities[alien.Location] = existing
	}

	return aliensByCities
}

func (e *Engine) removeAliens(aliens []*types.Alien) {
	for _, a := range aliens {
		for i, aa := range e.aliens {
			if a.Name == aa.Name {
				e.aliens = append(e.aliens[:i], e.aliens[i+1:]...)
			}
		}
	}
}

func getCasualties(aliensByCities map[types.City][]*types.Alien) ([]*types.Alien, []types.City) {
	var (
		destroyedCities []types.City
		deadAliens      []*types.Alien
	)

	for city, aliens := range aliensByCities {
		if len(aliens) > 1 {
			destroyedCities = append(destroyedCities, city)
			deadAliens = append(deadAliens, aliens...)

			displayDestroyed(city, aliens)
		}
	}

	return deadAliens, destroyedCities
}

func displayDestroyed(ciy types.City, aliens []*types.Alien) {
	var res string
	for _, alien := range aliens {
		res += alien.Name + " "
	}

	log.Println(ciy, "has been destroyed by aliens:", res)
}
